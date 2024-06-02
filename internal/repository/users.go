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
	Email    string    `json:"email" db:"email"`

	Password string `json:"password" db:"password"`

	FirstName string         `json:"first_name" db:"first_name"`
	LastName  sql.NullString `json:"last_name" db:"last_name"`
	JobTitle  sql.NullString `json:"job_title" db:"job_title"`

	Language sql.NullString `json:"language" db:"language"`
	Photo    sql.NullString `json:"photo" db:"photo"`

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
		"INSERT INTO users (id, username, first_name, email, password, created_at, last_name) values ($1, $2, $3, $4, $5, $6, $7);",
		user.ID, user.Username, user.FirstName, user.Email, user.Password, user.CreatedAt, user.LastName,
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
