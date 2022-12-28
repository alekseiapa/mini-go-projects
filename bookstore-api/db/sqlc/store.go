package db

import (
	"database/sql"
)

// Store provide all functions to execute db queries and transactions
// In order to make a support of transactions we should use the Composition here

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}
