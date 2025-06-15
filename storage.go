package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Storage interface {
	// GetAccount retrieves an account by ID.
	CreateAccount(*Account) (*Account, error)
	DeleteAccount(int) error
	UpdateAccount(*Account) (*Account, error)
	GetAccount(int) (*Account, error)
}

type PostgresStorage struct {
	// db *sql.DB // Assuming you have a database connection
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, error) {
	connStr := "user=postgres dbname=postgres password=go-bank host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStorage{db: db}, nil
}
