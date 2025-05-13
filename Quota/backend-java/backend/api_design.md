# Progettazione delle API per Quota

## **1. Autenticazione**
### Endpoint:
- **POST /api/auth/register**
  - Descrizione: Registra un nuovo utente.
  - Input:
    ```json
    {
      "email": "string",
      "password": "string",
      "name": "string"
    }
    ```
  - Output:
    ```json
    {
      "message": "Utente registrato con successo",
      "userId": "string"
    }
    ```

- **POST /api/auth/login**
  - Descrizione: Effettua il login di un utente.
  - Input:
    ```json
    {
      "email": "string",
      "password": "string"
    }
    ```
  - Output:
    ```json
    {
      "token": "string",
      "userId": "string"
    }
    ```

- **POST /api/auth/logout**
  - Descrizione: Effettua il logout dell'utente.
  - Input: Nessuno.
  - Output:
    ```json
    {
      "message": "Logout effettuato con successo"
    }
    ```

---

## **2. Gestione gruppi**
### Endpoint:
- **POST /api/groups**
  - Descrizione: Crea un nuovo gruppo.
  - Input:
    ```json
    {
      "name": "string",
      "members": ["userId1", "userId2"]
    }
    ```
  - Output:
    ```json
    {
      "groupId": "string",
      "message": "Gruppo creato con successo"
    }
    ```

- **GET /api/groups**
  - Descrizione: Recupera tutti i gruppi dell'utente autenticato.
  - Input: Nessuno.
  - Output:
    ```json
    [
      {
        "groupId": "string",
        "name": "string",
        "members": ["userId1", "userId2"]
      }
    ]
    ```

- **DELETE /api/groups/{groupId}**
  - Descrizione: Elimina un gruppo.
  - Input: Nessuno.
  - Output:
    ```json
    {
      "message": "Gruppo eliminato con successo"
    }
    ```

---

## **3. Gestione spese**
### Endpoint:
- **POST /api/expenses**
  - Descrizione: Aggiunge una nuova spesa a un gruppo.
  - Input:
    ```json
    {
      "groupId": "string",
      "amount": "number",
      "description": "string",
      "paidBy": "userId",
      "participants": ["userId1", "userId2"]
    }
    ```
  - Output:
    ```json
    {
      "expenseId": "string",
      "message": "Spesa aggiunta con successo"
    }
    ```

- **GET /api/expenses/{groupId}**
  - Descrizione: Recupera tutte le spese di un gruppo.
  - Input: Nessuno.
  - Output:
    ```json
    [
      {
        "expenseId": "string",
        "amount": "number",
        "description": "string",
        "paidBy": "userId",
        "participants": ["userId1", "userId2"],
        "date": "string"
      }
    ]
    ```

- **DELETE /api/expenses/{expenseId}**
  - Descrizione: Elimina una spesa.
  - Input: Nessuno.
  - Output:
    ```json
    {
      "message": "Spesa eliminata con successo"
    }
    ```

---

## **4. Eliminazione account**
### Endpoint:
- **DELETE /api/users/{userId}**
  - Descrizione: Elimina l'account dell'utente.
  - Input: Nessuno.
  - Output:
    ```json
    {
      "message": "Account eliminato con successo"
    }
    ```

---

## **5. Modifica gruppi**
### Endpoint:
- **PUT /api/groups/{groupId}**
  - Descrizione: Modifica i dettagli di un gruppo.
  - Input:
    ```json
    {
      "name": "string",
      "members": ["userId1", "userId2"]
    }
    ```
  - Output:
    ```json
    {
      "message": "Gruppo aggiornato con successo"
    }
    ```

---

## **6. Modifica spese**
### Endpoint:
- **PUT /api/expenses/{expenseId}**
  - Descrizione: Modifica i dettagli di una spesa.
  - Input:
    ```json
    {
      "amount": "number",
      "description": "string",
      "paidBy": "userId",
      "participants": ["userId1", "userId2"]
    }
    ```
  - Output:
    ```json
    {
      "message": "Spesa aggiornata con successo"
    }
    ```

---

## **7. Recupero dettagli utente**
### Endpoint:
- **GET /api/users/me**
  - Descrizione: Recupera i dettagli dell'utente autenticato.
  - Input: Nessuno.
  - Output:
    ```json
    {
      "userId": "string",
      "email": "string",
      "name": "string"
    }
    ```

---

## **8. Recupero saldi**
### Endpoint:
- **GET /api/groups/{groupId}/balances**
  - Descrizione: Recupera i saldi tra i membri di un gruppo.
  - Input: Nessuno.
  - Output:
    ```json
    [
      {
        "userId": "string",
        "balance": "number"
      }
    ]
    ```

---

## **9. Gestione notifiche**
### Endpoint:
- **GET /api/notifications**
  - Descrizione: Recupera le notifiche per l'utente autenticato.
  - Input: Nessuno.
  - Output:
    ```json
    [
      {
        "notificationId": "string",
        "message": "string",
        "date": "string"
      }
    ]
    ```
