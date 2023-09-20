package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type storage interface {
	CreateAccount(account *Account) error
	DeleteAccount(id int) error
	UpdateAccount(account *Account) error
	GetAccountById(id int) (*Account, error)
	GetAccounts() ([]*Account, error)
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

func (s *PostgresStore) CreateAccount(a *Account) error {
	query := `
insert into account(first_name, last_name, number, balance, created_at)
values ($1, $2, $3, $4, $5)
`
	rows, err := s.db.Query(query, a.FirstName, a.LastName, a.Number, a.Balance, a.CreatedAt)

	log.Printf("%+v\n", rows)
	return err
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

func (s *PostgresStore) GetAccounts() (a []*Account, err error) {
	rows, err := s.db.Query("select * from account")

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		if err := rows.Scan(
			&account.ID,
			&account.FirstName,
			&account.LastName,
			&account.Number,
			&account.Balance,
			&account.CreatedAt,
		); err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s *PostgresStore) Init() {
	err := s.createAccountTableForTest()
	if err != nil {
		panic(err)
	}
}

// This is only used for Test (You shouldn't implement create method like this below in your real project, it just an simple project to learn Go)
func (s *PostgresStore) createAccountTableForTest() error {
	query := `create table if not exists account(
    	id serial primary key,
    	first_name varchar(50),
    	last_name varchar(50),
    	number serial,
    	balance serial,
    	created_at timestamp
    )`

	log.Println("Creating the table for testing app...")

	// You should execute query with Context, do not like this
	if _, err := s.db.Exec(query); err != nil {
		return err
	}

	log.Println("Created the table for testing app")

	return nil
}
