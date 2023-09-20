package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type storage interface {
	CreateAccount(account *Account) error
	DeleteAccount(id int) error
	UpdateAccount(account *Account) error
	GetAccountById(id int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	// It should be changed to use a system environment
	connStr := "user=postgres dbname=postgres password=roach sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) CreateAccount(account *Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(account *Account) error {
	return nil
}

func (s *PostgresStore) GetAccountById(id int) (a *Account, err error) {
	return nil, nil
}

func (s *PostgresStore) Init() {
	s.createAccountTableForTest()
}

// This is only used for Test (You shouldn't implement create method like this below in your real project, it just an simple project to learn Go)
func (s *PostgresStore) createAccountTableForTest() error {
	query := `create table if not exists account(
    	id serial primary key,
    	first_name varchar(50),
    	last_name varchar(50),
    	number serial,
    	balance,
    	created_at timestamp
    )`

	// You should execute query with Context, do not like this
	if _, err := s.db.Exec(query); err != nil {
		return fmt.Errorf("your table didn't create, you should check query")
	}

	return nil
}
