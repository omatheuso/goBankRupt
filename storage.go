package main

import "database/sql"

type Storage interface {
	CreateAccount(account *Account) error
	DeleteAccount(id int) error
	UpdateAccount(account *Account) error
	GetAccount(id int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	return nil, nil
}
