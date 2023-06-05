package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"graded/model"
	"graded/pkg/erring"
)

type userRepository struct {
	db *sql.DB
}

func InitUserRepo(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

const userPath = `-> userRepository:%s(): %w`

func (r *userRepository) Create(ctx context.Context, user model.User) (int64, error) {
	query := `
	insert into users (login, email, hash_password)
	values ($1, $2, $3)`

	res, err := r.db.ExecContext(ctx, query, &user.Login, &user.Email, &user.Password)
	if err != nil {
		if isDuplicate(err) {
			return -1, fmt.Errorf(userPath, "Create", erring.ErrDuplicateValue)
		}
		return -1, fmt.Errorf(userPath, "Create", err)
	}

	uid, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf(userPath, "Create", err)
	}

	return uid, nil
}

func (r *userRepository) GetByID(ctx context.Context, uid int64) (user model.User, err error) {
	query := `
	select 
		id,
		login,
		email,
		hash_password,
		is_active,
		version
	from users
	where id = $1`

	if err := r.db.QueryRowContext(ctx, query, uid).Scan(
		&user.ID,
		&user.Login,
		&user.Email,
		&user.Password,
		&user.IsActive,
		&user.Version,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.User{}, fmt.Errorf(userPath, "GetByID", erring.ErrNoData)
		}
		return model.User{}, fmt.Errorf(userPath, "GetByID", err)
	}

	return user, nil
}
