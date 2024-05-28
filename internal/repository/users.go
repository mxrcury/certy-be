package repository

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/mxrcury/certy/pkg/logger"
)

type User struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Username string    `json:"username" db:"username"`
	Password string    `json:"password" db:"password"`
	Email    string    `json:"email" db:"email"`

	CreatedAt string         `json:"created_at" db:"created_at"`
	UpdatedAt sql.NullString `json:"updated_at" db:"updated_at"`
}

type UsersRepo struct {
	db *sqlx.DB
}

func NewUsersRepo(db *sqlx.DB) Users {
	return &UsersRepo{db: db}
}

func (u *UsersRepo) Create(user *User) error {
	if _, err := u.db.Exec(
		"INSERT INTO users (id, username, email, password, created_at) values ($1, $2, $3, $4, $5);",
		user.ID, user.Username, user.Email, user.Password, user.CreatedAt,
	); err != nil {
		return err
	}

	return nil
}

func (u *UsersRepo) Update(id uuid.UUID, user *User) (*User, error) {
	return nil, nil
}

func (u *UsersRepo) GetByEmailOrUsername(email string, username string) *User {
	var user User

	if err := u.db.Get(&user, "SELECT * FROM users WHERE email = $1 OR username = $2;", email, username); err != nil {
		logger.Error(err.Error())
		return nil
	}

	return &user
}

func (u *UsersRepo) GetByUsername(username string) *User {
	var user User

	if err := u.db.Get(&user, "SELECT * FROM users WHERE username = $1;", username); err != nil {
		return nil
	}

	return &user
}

func (u *UsersRepo) GetAll(pagination *Pagination) []User {
	return nil
}

func (u *UsersRepo) DeleteByID(id uuid.UUID) (*User, error) {
	return nil, nil
}
