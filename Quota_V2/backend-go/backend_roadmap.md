# Roadmap per il Backend di Quota

## **Fase 1: Setup del Progetto**
1. **Inizializzazione del progetto**:
   - ~~Creare una nuova directory per il backend (es. `backend-go`).~~ **Completato**
   - ~~Inizializzare un modulo Go:~~ **Completato**
     ```bash
     go mod init quota-backend
     ```
   - ~~Configurare un file `.env` per le variabili d'ambiente (es. connessione al database).~~ **Completato**

2. **Setup del database**:
   - ~~Configurare la connessione a PostgreSQL utilizzando una libreria come `pgx` o `gorm`.~~ **Completato**
   - ~~Creare un file di migrazione per sincronizzare il database con il codice.~~ **Non necessario**

3. **Setup del server**:
   - ~~Utilizzare un framework come **Gin** o **Echo** per gestire le API REST.~~ **Completato**
   - ~~Configurare il routing di base.~~ **Completato**

4. **Struttura del progetto**:
   - ~~Organizzare il codice in pacchetti:~~ **Completato**
     - `controllers`: Logica delle API.
     - `models`: Definizione delle entità e interazioni con il database.
     - `routes`: Definizione degli endpoint.
     - `utils`: Funzioni di utilità (es. hashing password, validazione input).

---

## **Fase 2: Sviluppo delle API**
### **Autenticazione**
- **POST /register**: ~~Registra un nuovo utente.~~ **Completato**
- **POST /login**: ~~Effettua il login e restituisce un token JWT.~~ **Completato**
- **GET /profile**: ~~Restituisce i dati dell'utente autenticato.~~ **In corso**
- **POST /logout**: **Da implementare**

### **Gestione Gruppi**
- **POST /groups**: **Da implementare**
- **GET /groups**: **Da implementare**
- **GET /groups/:id**: **Da implementare**
- **PUT /groups/:id**: **Da implementare**
- **DELETE /groups/:id**: **Da implementare**

### **Gestione Spese**
- **POST /expenses**: **Da implementare**
- **GET /expenses**: **Da implementare**
- **PUT /expenses/:id**: **Da implementare**
- **DELETE /expenses/:id**: **Da implementare**

### **Saldi**
- **GET /balances/:groupId**: **Da implementare**

### **Gestione Inviti ai Gruppi**
- **POST /groups/:id/invite**: **Da implementare**
- **POST /groups/:id/join**: **Da implementare**

### **Gestione Avatar Utente**
- **GET /avatars**: **Da implementare**
- **PUT /profile/avatar**: **Da implementare**

### **Log delle Modifiche**
- **GET /logs**: **Da implementare**

### **Notifiche**
- **POST /notifications**: **Da implementare**
- **GET /notifications**: **Da implementare**

### **Ricerca e Filtri**
- **GET /expenses/search**: **Da implementare**

### **Statistiche**
- **GET /stats**: **Da implementare**

---

## **Fase 3: Sicurezza, Prestazioni e Gestione degli Errori**
1. **Sicurezza**:
   - ~~Implementare l'autenticazione con JWT per proteggere le API.~~ **Completato**
   - ~~Validare tutti gli input per prevenire attacchi SQL injection, XSS e CSRF.~~ **Completato**
   - ~~Proteggere le API con rate limiting per prevenire abusi.~~ **Completato**
   - ~~Utilizzare HTTPS per tutte le comunicazioni.~~ **Completato**
   - ~~Configurare un sistema di logging per monitorare accessi sospetti e tentativi di attacco.~~ **Completato**

2. **Prestazioni**:
   - **Da implementare**

3. **Gestione degli Errori**:
   - ~~Implementare un middleware globale per intercettare e gestire gli errori.~~ **Completato**
   - ~~Restituire messaggi di errore standardizzati con codici di stato HTTP appropriati.~~ **Completato**
   - ~~Loggare tutti gli errori critici in un sistema centralizzato (es. Sentry o ELK Stack).~~ **Completato**
   - ~~Creare una documentazione per gli errori comuni e le relative soluzioni per il team di sviluppo.~~ **Completato**

---

## **Fase 4: Testing**
1. **Test delle API**:
   - **Da implementare**

2. **Test di carico**:
   - **Da implementare**

---

## **Fase 5: Deploy**
1. **Containerizzazione**:
   - **Da implementare**

2. **Deploy su cloud**:
   - **Da implementare**
