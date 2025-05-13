package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/sirupsen/logrus"
)

var jwtSecret = []byte("your_jwt_secret")

// GenerateJWT genera un token JWT per un utente
func GenerateJWT(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Scadenza di 24 ore
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateJWT valida un token JWT e restituisce i claims
func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}

// MiddlewareJWT verifica il token JWT e protegge le rotte
func MiddlewareJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			logrus.Error("Token mancante nell'header Authorization")
			c.JSON(401, gin.H{"error": "Token mancante"})
			c.Abort()
			return
		}

		claims, err := ValidateJWT(tokenString)
		if err != nil {
			logrus.WithError(err).Error("Errore durante la validazione del token JWT")
			c.JSON(401, gin.H{"error": "Token non valido"})
			c.Abort()
			return
		}

		logrus.WithField("claims", claims).Info("Claims estratti dal token JWT")

		// Verifica che il campo `user_id` sia presente nei claims
		userID, ok := claims["user_id"].(float64)
		if !ok {
			logrus.Error("Campo user_id mancante o non valido nei claims del token JWT")
			c.JSON(401, gin.H{"error": "Token non valido: ID utente mancante"})
			c.Abort()
			return
		}

		// Aggiungi l'ID utente al contesto
		logrus.WithField("user_id", userID).Info("ID utente aggiunto al contesto")
		c.Set("user_id", int64(userID))
		c.Next()
	}
}
