package main

import (
	"context"
	"io"
	"log"
	"os"
	"path/filepath"

	"Quota_V2/backend-go/routes"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	// Creazione della cartella per i log
	logDir := "logs"
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		logrus.Fatalf("Errore durante la creazione della cartella dei log: %v", err)
	}

	// Configurazione di Logrus per scrivere sia su file che in console
	logFilePath := filepath.Join(logDir, "backend.log")
	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		logrus.Fatalf("Errore durante l'apertura del file di log: %v", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(multiWriter)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

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

	// Configurare le rotte
	routes.SetupRoutes(r)

	// Porta del server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta di default
	}

	r.Run(":" + port) // Avvia il server
}
