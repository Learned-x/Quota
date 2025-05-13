package controllers

import (
	"Quota_V2/backend-go/models"
	"Quota_V2/backend-go/utils"
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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
		logrus.WithFields(logrus.Fields{
			"email":  requestBody.Email,
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		}).Error("Errore durante il binding dei dati JSON")
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
		logrus.WithFields(logrus.Fields{
			"email":  requestBody.Email,
			"status": http.StatusBadRequest,
			"error":  err.Error(),
		}).Error("Errore durante il binding dei dati JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dati non validi"})
		return
	}

	// Recupera l'utente dal database
	user, err := models.GetUserByEmail(context.Background(), requestBody.Email)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"email":  requestBody.Email,
			"status": http.StatusUnauthorized,
			"error":  err.Error(),
		}).Error("Errore durante il recupero dell'utente")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenziali non valide"})
		return
	}

	logrus.WithFields(logrus.Fields{
		"user_id":  user.ID,
		"username": user.Username,
		"email":    user.Email,
	}).Info("Utente trovato")

	// Confronta la password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(requestBody.Password)); err != nil {
		logrus.WithFields(logrus.Fields{
			"email":  requestBody.Email,
			"status": http.StatusUnauthorized,
		}).Error("Password errata")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenziali non valide"})
		return
	}

	// Genera il token JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"user_id": user.ID,
			"status":  http.StatusInternalServerError,
			"error":   err.Error(),
		}).Error("Errore durante la generazione del token JWT")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore interno del server"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// GetProfile restituisce i dati completi dell'utente autenticato
func GetProfile(c *gin.Context) {
	userID := c.GetString("user_id")
	logrus.WithField("user_id", userID).Info("Recupero del profilo utente iniziato")

	if userID == "" {
		logrus.Error("ID utente mancante nel contesto")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "ID utente mancante"})
		return
	}

	// Recupera i dati dell'utente dal database
	user, err := models.GetUserByID(context.Background(), userID)
	if err != nil {
		logrus.WithField("user_id", userID).WithError(err).Error("Errore durante il recupero del profilo utente")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore durante il recupero del profilo utente"})
		return
	}

	logrus.WithField("user_id", userID).Info("Profilo utente recuperato con successo")
	c.JSON(http.StatusOK, gin.H{
		"id":         user.ID,
		"username":   user.Username,
		"email":      user.Email,
		"nome":       user.Nome,
		"cognome":    user.Cognome,
		"created_at": user.CreatedAt,
		"id_avatar":  user.IDAvatar,
	})
}

// Logout invalida il token JWT (placeholder per ora)
func Logout(c *gin.Context) {
	logrus.Info("Logout eseguito con successo")
	c.JSON(http.StatusOK, gin.H{"message": "Logout eseguito con successo"})
}
