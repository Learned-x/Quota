# Descrizioni dettagliate delle schermate principali per Quota

## **Linee guida per il design**
### Design minimal:
- Utilizzare un layout semplice e intuitivo.
- Evitare elementi visivi superflui.
- Garantire una navigazione chiara e coerente.

### Colori di default:
- **Primario**: #4CAF50 (verde chiaro per pulsanti e accenti principali).
- **Secondario**: #FFFFFF (bianco per sfondi principali).
- **Testo**: #212121 (nero per il testo principale).
- **Sfondo**: #F5F5F5 (grigio chiaro per sfondi secondari).
- **Errore**: #F44336 (rosso per messaggi di errore).

## **1. Onboarding: Registrazione e Login**
### Descrizione:
Questa schermata consente agli utenti di registrarsi o accedere all'applicazione. Deve essere semplice e intuitiva, con opzioni per il login tramite email/password o provider esterni (Google, Apple).

### Elementi principali:
- **Logo dell'app**: Posizionato in alto al centro.
- **Form di login**:
  - Campo email.
  - Campo password.
  - Pulsante "Accedi".
- **Link per il recupero password**: "Hai dimenticato la password?".
- **Pulsante di registrazione**: "Non hai un account? Registrati".
- **Opzioni di login social**:
  - Pulsanti per Google e Apple.
- **Messaggi di errore**: Mostrati in caso di credenziali errate o campi mancanti.

---

## **2. Dashboard: Elenco gruppi e riepilogo spese**
### Descrizione:
La dashboard è la schermata principale dopo il login. Mostra un elenco dei gruppi a cui l'utente appartiene e un riepilogo delle spese recenti.

### Elementi principali:
- **Barra di navigazione**:
  - Icona "Home".
  - Icona "Profilo".
  - Icona "Impostazioni".
- **Elenco gruppi**:
  - Nome del gruppo.
  - Totale spese del gruppo.
  - Pulsante per accedere ai dettagli del gruppo.
- **Riepilogo spese recenti**:
  - Ultime 5 spese con importo, categoria e data.
- **Pulsante "Crea nuovo gruppo"**: Posizionato in basso a destra.

---

## **3. Dettagli del gruppo: Elenco spese e saldi**
### Descrizione:
Questa schermata mostra i dettagli di un gruppo specifico, inclusi l'elenco delle spese e i saldi tra i membri.

### Elementi principali:
- **Nome del gruppo**: In alto al centro.
- **Elenco spese**:
  - Nome della spesa.
  - Importo.
  - Data.
  - Chi ha pagato.
  - Pulsanti per modificare o eliminare la spesa.
- **Sezione saldi**:
  - Elenco dei debiti/crediti tra i membri.
  - Totale saldo per ogni membro.
- **Pulsante "Aggiungi spesa"**: Posizionato in basso a destra.

---

## **4. Aggiunta spesa: Form per nuova spesa**
### Descrizione:
Questa schermata consente agli utenti di aggiungere una nuova spesa a un gruppo.

### Elementi principali:
- **Form di inserimento**:
  - Campo "Nome spesa".
  - Campo "Importo".
  - Campo "Categoria" (menu a tendina).
  - Campo "Data" (selezione calendario).
  - Campo "Chi ha pagato" (menu a tendina con i membri del gruppo).
  - Selezione "Chi partecipa" (checkbox per i membri del gruppo).
- **Pulsante "Salva"**: Per confermare l'aggiunta della spesa.
- **Pulsante "Annulla"**: Per tornare indietro senza salvare.

---

## **5. Invito al gruppo: Generazione e condivisione di link**
### Descrizione:
Questa schermata consente agli utenti di generare un link di invito per aggiungere nuovi membri a un gruppo.

### Elementi principali:
- **Nome del gruppo**: In alto al centro.
- **Sezione link di invito**:
  - Campo con il link generato.
  - Pulsante "Copia link".
  - Pulsante "Condividi" (integrazione con app di messaggistica).
- **Elenco membri attuali**:
  - Nome e avatar di ogni membro.
  - Ruolo (es. amministratore, membro).

---

## **6. Recupero password: Reimposta credenziali dimenticate**
### Descrizione:
Questa schermata consente agli utenti di reimpostare la password in caso di smarrimento.

### Elementi principali:
- **Form di recupero**:
  - Campo email.
  - Pulsante "Invia email di recupero".
- **Messaggi di conferma**: Mostrati dopo l'invio dell'email.
- **Pulsante "Torna al login"**: Per tornare alla schermata di accesso.
