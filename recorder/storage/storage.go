package storage

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/q-sw/go-bank-analysis/types"
	"log"
	"os"
	"time"
)

type Storage interface {
	AddBalance() error
	CreateTransaction() error
	UpdateTransaction() error
	GetTransaction() error
	GetAllTransaction() error
}

type PostgreSQL struct {
	db *sql.DB
}

func NewConnection() (*PostgreSQL, error) {
	db_user := os.Getenv("DB_USER")
	db_pwd := os.Getenv("DB_PWD")
	db_name := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("user=%v dbname=%v password=%v host='172.17.0.2' port='5432' sslmode=disable", db_user, db_name, db_pwd)
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	fmt.Println("Connect")
	return &PostgreSQL{db: db}, nil
}

func (s *PostgreSQL) Init() error {
	c := s.createCategoriesTable()
	if c != nil {
		return c
	}

	t := s.createTransactionTable()
	if t != nil {
		return t
	}

	b := s.createBalanceTable()
	if b != nil {
		return b
	}

	return nil
}

func (s *PostgreSQL) createTransactionTable() error {
	query := `create table if not exists transactions (
        id int not null primary key,
        transction_id varchar(250),
        amount varchar(250),
        code varchar(250),
        description varchar(250),
        categoryID int references categories(id)
    )`

	_, err := s.db.Exec(query)
	return err

}

func (s *PostgreSQL) createCategoriesTable() error {
	query := `create table if not exists categories (
        id int not null primary key,
        name varchar(250)
    )`

	_, err := s.db.Exec(query)
	return err

}
func (s *PostgreSQL) createBalanceTable() error {
	query := `create table if not exists balance (
        id serial primary key,
        date timestamp,
        amount decimal
    )`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgreSQL) AddBalance(bal *types.Balance) error {
	log.Println("add balance")
	r, err := s.db.Prepare("insert into balance (date, amount) values ($1, $2)")
	if err != nil {
		return err
	}
	defer r.Close()

	_, err = r.Exec(time.Now(), bal.TransctionAmount.Amount)

	if err != nil {
		return err
	}
	return nil
}
