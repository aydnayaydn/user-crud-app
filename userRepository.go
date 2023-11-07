package main

import (
	"database/sql"
	"log"
)

type UserRepository interface {
	CreateUser(User User) error
	GetUser(id int) (User, error)
	UpdateUser(id int, User User) error
	DeleteUser(id int) error
}

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{db: db}
}

func (r *SQLUserRepository) CreateUser(User User) error {
	query := "INSERT User (name, surname) VALUES (?, ?)"
	_, err := r.db.Exec(query, User.Name, User.Surname)
	return err
}

func (r *SQLUserRepository) GetUser(id int) (User, error) {
	var User User
	query := "SELECT id, name, surname FROM User WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&User.ID, &User.Name)
	return User, err
}

func (r *SQLUserRepository) UpdateUser(id int, User User) error {
	query := "UPDATE User SET name = ?, surname = ? WHERE id = ?"
	_, err := r.db.Exec(query, User.Name, User.Surname, id)
	return err
}

func (r *SQLUserRepository) DeleteUser(id int) error {
	query := "DELETE FROM User WHERE id = ?"
	_, err := r.db.Exec(query, id)
	return err
}

func InitDBConnection() (*sql.DB, error) {
	connString := "server=server_address;user id=user_name;password=password;database=database_name"
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
