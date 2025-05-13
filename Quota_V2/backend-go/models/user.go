package models

import (
	"context"
	"errors"

	"Quota_V2/backend-go/db"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type User struct {
	ID       int64
	Username string
	Email    string
	Password string
	Nome     string
	Cognome  string
}

// CreateUser salva un nuovo utente nel database
func CreateUser(ctx context.Context, user User) error {
	conn, err := connectToDatabase()
	if err != nil {
		logrus.WithError(err).Error("Errore durante la connessione al database")
		return err
	}
	defer conn.Close(ctx)

	query := `INSERT INTO utente (username, email, password, nome, cognome, provider, created_at) VALUES ($1, $2, $3, $4, $5, $6, NOW())`
	_, err = conn.Exec(ctx, query, user.Username, user.Email, user.Password, user.Nome, user.Cognome, "email")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"query":    query,
			"username": user.Username,
			"email":    user.Email,
		}).WithError(err).Error("Errore durante l'inserimento dell'utente")
		return errors.New("errore durante l'inserimento dell'utente")
	}

	logrus.WithFields(logrus.Fields{
		"username": user.Username,
		"email":    user.Email,
	}).Info("Utente creato con successo")

	return nil
}

// GetUserByEmail recupera un utente dal database tramite email
func GetUserByEmail(ctx context.Context, email string) (User, error) {
	conn, err := connectToDatabase()
	if err != nil {
		return User{}, err
	}
	defer conn.Close(ctx)

	var user User
	query := `SELECT id, username, email, password, nome, cognome FROM utenti WHERE email = $1`
	row := conn.QueryRow(ctx, query, email)
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Nome, &user.Cognome); err != nil {
		return User{}, errors.New("utente non trovato")
	}

	return user, nil
}

func connectToDatabase() (*pgx.Conn, error) {
	return db.ConnectToDatabase()
}
