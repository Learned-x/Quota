package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func init() {
	// Determina l'ambiente corrente
	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "development" // Ambiente di default
	}

	// Carica il file .env appropriato
	envFile := ".env." + appEnv
	if err := godotenv.Load(envFile); err != nil {
		log.Printf("Attenzione: impossibile caricare il file %s", envFile)
	}
}

func connectToDatabase() *pgx.Conn {
	// Ottieni le variabili d'ambiente per la connessione al database
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL non è configurato")
	}

	// Connessione al database
	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Impossibile connettersi al database: %v", err)
	}

	log.Println("Connessione al database riuscita!")
	return conn
}

func main() {
	// Connessione al database
	dbConn := connectToDatabase()
	defer dbConn.Close(context.Background())

	r := gin.Default()

	// Endpoint di test
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Porta del server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta di default
	}

	r.Run(":" + port) // Avvia il server
}
