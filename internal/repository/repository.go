package repository

import (
	"awesomeProject/internal/entity"
	"errors"
	"github.com/jackc/pgx"
)

type Repository interface {
	CreateUser(entity.User) error
	GetUser(id string) (entity.User, error)
	UpdateUser(entity.User) (entity.User, error)
	DeleteUser(id string) error
}

type repo struct {
	conn *pgx.Conn
}

func NewRepository(conn *pgx.Conn) Repository {
	return &repo{conn: conn}
}
func (r *repo) CreateUser(user entity.User) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("INSERT INTO Users VALUES ($1, $2, $3)", user.NickName, user.Email, user.Password)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
func (r *repo) GetUser(id string) (entity.User, error) {
	var user entity.User
	rows, _ := r.conn.Query("SELECT * FROM Users WHERE id = $1", id)
	if rows.Next() {
		if rows.Err() != nil {
			return entity.User{}, rows.Err()
		}
		err := rows.Scan(&user.NickName, &user.Email, &user.Password)
		if err != nil {
			return entity.User{}, err
		}
	}
	if user.Email == "" {
		return entity.User{}, errors.New("user not found")
	}
	return user, nil
}
func (r *repo) UpdateUser(user entity.User) (entity.User, error) {
	tx, err := r.conn.Begin()
	if err != nil {
		return entity.User{}, err
	}
	_, err = tx.Exec("UPDATE Users SET nickname = $2, email = $3, password = $4 WHERE id = $1",
		user.ID, user.NickName, user.Email, user.Password)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return entity.User{}, err
		}
	}
	err = tx.Commit()
	if err != nil {
		return entity.User{}, err
	}
	return r.GetUser(user.ID)
}
func (r *repo) DeleteUser(id string) error {
	tx, err := r.conn.Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("DELETE FROM Users WHERE id = $1", id)
	if err != nil {
		err = tx.Rollback()
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
