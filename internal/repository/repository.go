package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repositories struct {
	UsersRepository Users
}

type (
	Users interface {
		Create(user *User) error
		Update(id uuid.UUID, user *User) (*User, error)
		GetByEmailOrUsername(email string, username string) *User
		GetByEmail(email string) *User
		GetAll(pagination *Pagination) []User
		DeleteByID(id uuid.UUID) (*User, error)
	}
)

type (
	Pagination struct {
		Page int
		Size int
	}
)

func NewRepositories(db *sqlx.DB) *Repositories {
	usersRepository := NewUsersRepo(db)

	return &Repositories{UsersRepository: usersRepository}
}
