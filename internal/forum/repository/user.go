package repository

import (
	"context"
	"database/sql"
	"forum/internal/forum/model"
	"time"
)

// type UserStorage interface {
// 	Insert(user *model.CreateUserDTO) (uint64, error)
// 	GetUserByID(id uint64) (*model.User, error)
// 	GetUserByEmail(email string) (*model.User, error)
// }

type userRepository struct {
	db *sql.DB
}

func NewRepositoryUser(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Insert(user *model.CreateUserDTO) (uint64, error) {
	stmt := `INSERT INTO users (name, email, hashpass) VALUES (?, ?, ?)`
	args := []interface{}{user.Name, user.Email, user.Password}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := u.db.ExecContext(ctx, stmt, args)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// if err == ErrNoRows -> return nil, nil
func (u *userRepository) GetUserByID(id uint64) (user *model.User, err error) {
	stmt := `SELECT * FROM users WHERE id == ?`
	args := []interface{}{id}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = u.db.QueryRowContext(ctx, stmt, args).Scan(user); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}

// if err == ErrNoRows -> return nil, nil
func (u *userRepository) GetUserByEmail(email string) (user *model.User, err error) {
	stmt := `SELECT * FROM users WHERE email == ?`
	args := []interface{}{email}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = u.db.QueryRowContext(ctx, stmt, args).Scan(user); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return user, nil
}
