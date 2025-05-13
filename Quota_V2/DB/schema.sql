-- Schema SQL per Quota: Applicazione di Gestione delle Spese Condivise

-- Creazione della tabella Avatar
CREATE TABLE Avatar (
    id SERIAL PRIMARY KEY,
    immagine_base64 TEXT NOT NULL, -- Salva l'immagine come stringa Base64
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Creazione della tabella Utenti
CREATE TABLE Utenti (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    cognome VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    provider VARCHAR(20) NOT NULL CHECK (provider IN ('Google', 'Apple', 'email')),
    id_avatar INT REFERENCES Avatar(id) ON DELETE SET NULL, -- Relazione con Avatar
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Creazione della tabella Gruppi
CREATE TABLE Gruppi (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    link_invito TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Creazione della tabella Spese
CREATE TABLE Spese (
    id SERIAL PRIMARY KEY,
    importo DECIMAL(10, 2) NOT NULL,
    categoria VARCHAR(50) NOT NULL,
    data DATE NOT NULL,
    id_pagante INT NOT NULL REFERENCES Utenti(id) ON DELETE CASCADE,
    id_gruppo INT NOT NULL REFERENCES Gruppi(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Creazione della tabella Partecipanti_Spese
CREATE TABLE Partecipanti_Spese (
    id SERIAL PRIMARY KEY,
    id_spesa INT NOT NULL REFERENCES Spese(id) ON DELETE CASCADE,
    id_partecipante INT NOT NULL REFERENCES Utenti(id) ON DELETE CASCADE,
    quota DECIMAL(10, 2) NOT NULL
);

-- Creazione della tabella Saldi
CREATE TABLE Saldi (
    id SERIAL PRIMARY KEY,
    id_creditore INT NOT NULL REFERENCES Utenti(id) ON DELETE CASCADE,
    id_debitore INT NOT NULL REFERENCES Utenti(id) ON DELETE CASCADE,
    id_gruppo INT NOT NULL REFERENCES Gruppi(id) ON DELETE CASCADE,
    saldo DECIMAL(10, 2) NOT NULL,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Funzione per aggiornare automaticamente i saldi
CREATE OR REPLACE FUNCTION aggiorna_saldi() RETURNS TRIGGER AS $$
BEGIN
    UPDATE Saldi
    SET saldo = saldo + NEW.importo / (SELECT COUNT(*) FROM Partecipanti_Spese WHERE id_spesa = NEW.id)
    WHERE id_creditore = NEW.id_pagante AND id_gruppo = NEW.id_gruppo;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger per la funzione aggiorna_saldi
CREATE TRIGGER trigger_aggiorna_saldi
AFTER INSERT OR UPDATE ON Spese
FOR EACH ROW
EXECUTE FUNCTION aggiorna_saldi();

-- Vista per riepilogare le spese per gruppo
CREATE VIEW Riepilogo_Spese AS
SELECT 
    g.nome AS nome_gruppo,
    SUM(s.importo) AS totale_spese,
    COUNT(s.id) AS numero_spese
FROM Gruppi g
JOIN Spese s ON g.id = s.id_gruppo
GROUP BY g.nome;

-- Vista per mostrare i saldi dettagliati
CREATE VIEW Dettagli_Saldi AS
SELECT 
    u1.nome AS creditore,
    u2.nome AS debitore,
    g.nome AS gruppo,
    s.saldo
FROM Saldi s
JOIN Utenti u1 ON s.id_creditore = u1.id
JOIN Utenti u2 ON s.id_debitore = u2.id
JOIN Gruppi g ON s.id_gruppo = g.id;

-- Creazione della tabella Log_Modifiche
CREATE TABLE Log_Modifiche (
    id SERIAL PRIMARY KEY,
    tabella VARCHAR(50) NOT NULL, -- Nome della tabella modificata
    operazione VARCHAR(10) NOT NULL CHECK (operazione IN ('INSERT', 'UPDATE', 'DELETE')),
    id_record INT NOT NULL, -- ID del record modificato
    dettagli JSONB, -- Dettagli della modifica in formato JSON
    timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Funzione per registrare le modifiche
CREATE OR REPLACE FUNCTION registra_modifica() RETURNS TRIGGER AS $$
BEGIN
    IF (TG_OP = 'DELETE') THEN
        INSERT INTO Log_Modifiche (tabella, operazione, id_record, dettagli)
        VALUES (TG_TABLE_NAME, TG_OP, OLD.id, row_to_json(OLD));
    ELSE
        INSERT INTO Log_Modifiche (tabella, operazione, id_record, dettagli)
        VALUES (TG_TABLE_NAME, TG_OP, NEW.id, row_to_json(NEW));
    END IF;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;

-- Trigger per la tabella Utenti
CREATE TRIGGER trigger_log_utenti
AFTER INSERT OR UPDATE OR DELETE ON Utenti
FOR EACH ROW
EXECUTE PROCEDURE registra_modifica();

-- Trigger per la tabella Gruppi
CREATE TRIGGER trigger_log_gruppi
AFTER INSERT OR UPDATE OR DELETE ON Gruppi
FOR EACH ROW
EXECUTE PROCEDURE registra_modifica();

-- Trigger per la tabella Spese
CREATE TRIGGER trigger_log_spese
AFTER INSERT OR UPDATE OR DELETE ON Spese
FOR EACH ROW
EXECUTE PROCEDURE registra_modifica();
