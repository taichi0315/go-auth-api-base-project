package persistence

import (
	"database/sql"

	"github.com/nepp-tumsat/documents-api/infrastructure/model"
	"golang.org/x/xerrors"
)

type UserRepository interface {
	SelectAll() ([]model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserDB(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) SelectAll() ([]model.User, error) {
	rows, err := u.db.Query(`
    SELECT username FROM users;
	`)

	if err != nil {
		err = xerrors.Errorf("Error in sql.DB: %v", err)
		return nil, err
	}

	var users []model.User
	var user model.User
	for rows.Next() {
		err := rows.Scan(&user.UserName)
		if err != nil {
			err = xerrors.Errorf("Error in sql.DB: %v", err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
