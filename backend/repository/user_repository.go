package repository

import (
	"backend/model"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	CreateUser(newUser *model.User) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository {
		db: db,
	}
}

func (repo *userRepository) CreateUser(newUser *model.User) (*model.User, error) {
	stmt, err := repo.db.Prepare("insert into users (id, username, email) values ($1,$2,$3) returning id")
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	defer stmt.Close()

	err = stmt.QueryRow(newUser.ID, newUser.UserName, newUser.Email).Scan(&newUser.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return newUser, nil
}