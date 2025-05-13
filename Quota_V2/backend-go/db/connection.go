package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

// ConnectToDatabase gestisce la connessione al database
func ConnectToDatabase() (*pgx.Conn, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL non è configurato")
	}

	conn, err := pgx.Connect(context.Background(), dbURL)
	if err != nil {
		log.Printf("Errore durante la connessione al database: %v", err)
		return nil, err
	}

	log.Println("Connessione al database riuscita!")
	return conn, nil
}
