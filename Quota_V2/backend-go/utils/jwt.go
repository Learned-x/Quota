package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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
			c.JSON(401, gin.H{"error": "Token mancante"})
			c.Abort()
			return
		}

		claims, err := ValidateJWT(tokenString)
		if err != nil {
			c.JSON(401, gin.H{"error": "Token non valido"})
			c.Abort()
			return
		}

		// Aggiungi i dati dell'utente al contesto
		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}
