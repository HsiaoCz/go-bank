package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
}

type MysqlStore struct {
	db *sql.DB
}

func NewPostgresStore() (*MysqlStore, error) {
	connStr := "root:root@tcp(127.0.0.1:3306)/go-bank"
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &MysqlStore{db: db}, nil
}

func (s *MysqlStore) CreateAccount(account *Account) error {
	sqlStr := `insert into account(first_name,last_name) values (?,?)`
	rest, err := s.db.Exec(sqlStr, account.FirstName, account.LastName)
	if err != nil {
		log.Fatal(err)
	}
	theId, err := rest.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(theId)
	return nil
}
func (s *MysqlStore) UpdateAccount(*Account) error {
	return nil
}
func (s *MysqlStore) DeleteAccount(id int) error {
	return nil
}
func (s *MysqlStore) GetAccountByID(id int) (*Account, error) {
	return nil, nil
}
