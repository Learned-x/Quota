# Roadmap per il Backend di Quota

## **Fase 1: Configurazione Iniziale**
1. **Setup del progetto**:
   - Configurare il progetto Spring Boot con Maven.
     - Creare un progetto su [Spring Initializr](https://start.spring.io/).
     - Aggiungere dipendenze: Spring Web, Spring Data JPA, Spring Security, PostgreSQL.
   - Verificare che il progetto si avvii correttamente con un endpoint di test.
2. **Configurazione del database**:
   - Configurare il file `application.properties`:
     - URL del database PostgreSQL.
     - Credenziali di accesso.
     - Dialetto Hibernate.
   - Creare le entità JPA basate sullo schema SQL:
     - `Avatar`, `Utente`, `Gruppo`, `Spesa`, `PartecipanteSpesa`.
     - Aggiungere annotazioni JPA per relazioni e vincoli.
3. **Versionamento**:
   - Inizializzare un repository Git:
     - Creare un commit iniziale con il setup del progetto.
   - Creare i branch principali:
     - `main` per il codice stabile.
     - `develop` per il codice in sviluppo.

---

## **Fase 2: Sviluppo delle API**
1. **Repository JPA**:
   - **Completato**
2. **Servizi**:
   - Implementare la logica di business:
     - **Autenticazione**:
       - Registrazione: Validare input, crittografare password con bcrypt.
       - Login: Generare token JWT.
       - Logout: Invalidare il token JWT.
     - **Gestione gruppi**:
       - Creazione: Validare input, generare link di invito univoco.
       - Modifica: Aggiornare nome o membri del gruppo.
       - Eliminazione: Rimuovere gruppo e spese associate.
       - Recupero: Restituire gruppi dell'utente autenticato.
     - **Gestione spese**:
       - Aggiunta: Validare input, calcolare quote per partecipanti.
       - Modifica: Aggiornare importo, descrizione o partecipanti.
       - Eliminazione: Rimuovere spesa e aggiornare saldi.
       - Recupero: Restituire spese di un gruppo.
     - **Calcolo saldi**:
       - Implementare logica per aggiornare saldi tra membri.
3. **Controller**:
   - Creare i controller per esporre le API definite nel file `api_design.md`:
     - `AuthController` per autenticazione.
     - `GruppoController` per gestione gruppi.
     - `SpesaController` per gestione spese.
     - `SaldoController` per calcolo saldi.

---

## **Fase 3: Sicurezza e Prestazioni**
1. **Spring Security**:
   - Configurare l'autenticazione basata su JWT:
     - Generare e validare token JWT.
     - Proteggere endpoint con autorizzazioni basate sui ruoli.
   - Implementare rate limiting per prevenire attacchi brute force.
2. **Caching**:
   - Integrare Redis per:
     - Cache delle risposte API frequenti.
     - Cache per riepiloghi delle spese e dettagli dei gruppi.
3. **Monitoraggio**:
   - Configurare strumenti come New Relic o Datadog:
     - Monitorare tempi di risposta delle API.
     - Configurare alert per anomalie nelle prestazioni.

---

## **Fase 4: Test e Documentazione**
1. **Test**:
   - Scrivere test unitari per:
     - Repository: Verificare query personalizzate.
     - Servizi: Validare logica di business.
     - Controller: Testare endpoint API.
   - Scrivere test di integrazione per:
     - Flussi completi (es. registrazione utente, creazione gruppo, aggiunta spesa).
2. **Documentazione**:
   - Integrare Swagger:
     - Generare documentazione automatica delle API.
     - Aggiungere descrizioni e esempi per ogni endpoint.

---

## **Fase 5: Rilascio**
1. **Preparazione**:
   - Creare un branch di rilascio (es. `release/1.0.0`).
   - Correggere eventuali bug e aggiornare la documentazione.
2. **Rilascio**:
   - Fare il merge su `main` e creare un tag per la versione (es. `v1.0.0`).
   - Distribuire il backend in produzione:
     - Configurare il server di produzione.
     - Verificare che tutte le funzionalità siano operative.

---

## **Aggiornamenti Recenti**

### **Fase 2: Sviluppo delle API**
- **Autenticazione**:
  - **Completato**: Registrazione con validazione input, crittografia password con bcrypt e valori predefiniti per `provider` e `username`.
  - **Completato**: Login con generazione di token JWT.
  - **Completato**: Logout con invalidazione dei token JWT e gestione degli errori per token scaduti o non validi.

### **Fase 4: Test e Documentazione**
- **Test di integrazione**:
  - **Completato**: Test per registrazione utenti e connessione al database PostgreSQL.
  - **Completato**: Test per il metodo `logoutUser` in scenari di token valido, scaduto e non valido.

### **Database**
- Rimosse tabelle ridondanti (`gruppo`, `utenti`, `spese`, `partecipanti_spese`) per semplificare la struttura del database.
- Aggiornata la struttura del database per mantenere solo le tabelle necessarie (`avatar`, `gruppi`, `utente`, `spesa`, `partecipante_spesa`, `saldi`).
