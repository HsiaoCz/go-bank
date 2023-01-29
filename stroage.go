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
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
}

type MysqlStore struct {
	db *sql.DB
}

func NewMysqlStore() (*MysqlStore, error) {
	connStr := "root:root@tcp(127.0.0.1:3306)/go-bank?charset=utf8mb4&parseTime=True&loc=Local"
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
	sqlStr := `insert into account(first_name,last_name,number,balance,created_at) values (?,?,?,?,?)`
	rest, err := s.db.Exec(sqlStr, account.FirstName, account.LastName, account.Number, account.Balance, account.CreatedAt)
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
func (s *MysqlStore) GetAccounts() ([]*Account, error) {
	sqlStr := `select id,first_name,last_name,number,balance,created_at from account where id>?`
	res, err := s.db.Query(sqlStr, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()
	account := make([]*Account, 0)
	for res.Next() {
		acc := Account{}
		err = res.Scan(&acc.ID, &acc.FirstName, &acc.LastName, &acc.Number, &acc.Balance, &acc.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		account = append(account, &acc)
	}
	return account, nil
}
