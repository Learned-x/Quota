package controllers

import (
	"Quota_V2/backend-go/models"
	"Quota_V2/backend-go/utils"
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterUser gestisce la registrazione di un nuovo utente
func RegisterUser(c *gin.Context) {
	// Estrarre i dati dal corpo della richiesta
	var requestBody struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Nome     string `json:"nome"`
		Cognome  string `json:"cognome"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dati non validi"})
		return
	}

	// Validare i dati di input
	if requestBody.Username == "" || requestBody.Email == "" || requestBody.Password == "" || requestBody.Nome == "" || requestBody.Cognome == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tutti i campi sono obbligatori"})
		return
	}

	// Hash della password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Errore durante l'hashing della password: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore interno del server"})
		return
	}

	// Salvataggio nel database
	newUser := models.User{
		Username: requestBody.Username,
		Email:    requestBody.Email,
		Password: string(hashedPassword),
		Nome:     requestBody.Nome,
		Cognome:  requestBody.Cognome,
	}

	if err := models.CreateUser(context.Background(), newUser); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			if strings.Contains(err.Error(), "ukgxvq4mjswnupehxnp35vawmo2") {
				c.JSON(http.StatusConflict, gin.H{"error": "Email già registrata"})
			} else if strings.Contains(err.Error(), "uk2vq82crxh3p7upassu0k1kmte") {
				c.JSON(http.StatusConflict, gin.H{"error": "Username già in uso"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore interno del server"})
			}
			return
		}

		log.Printf("Errore durante il salvataggio dell'utente: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore interno del server"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Utente registrato con successo"})
}

// Login gestisce l'autenticazione di un utente
func Login(c *gin.Context) {
	var requestBody struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dati non validi"})
		return
	}

	// Recupera l'utente dal database
	user, err := models.GetUserByEmail(context.Background(), requestBody.Email)
	if err != nil {
		log.Printf("Errore durante il recupero dell'utente: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenziali non valide"})
		return
	}

	// Confronta la password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenziali non valide"})
		return
	}

	// Genera il token JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		log.Printf("Errore durante la generazione del token JWT: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore interno del server"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
