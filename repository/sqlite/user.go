package sqlite

import (
	"context"
	"database/sql"
	"fmt"

	"graded/logger"
	"graded/model"
	"graded/pkg/errors"
)

type userRepository struct {
	logger *logger.Logger
	db     *sql.DB
}

func InitUserRepo(db *sql.DB, lg *logger.Logger) *userRepository {
	return &userRepository{
		db:     db,
		logger: lg,
	}
}

const userPath = `-> userRepository:%s(): %w`

func (r *userRepository) Create(ctx context.Context, user model.User) (int64, error) {
	query := `
	insert into users (login, email, hash_password)
	values ($1, $2, $3)`

	res, err := r.db.ExecContext(ctx, query, &user.Login, &user.Email, &user.Password)
	if isDuplicate(err) {
		return -1, fmt.Errorf(userPath, "Create", errors.ErrDuplicateValue)
	} else if err != nil {
		return -1, fmt.Errorf(userPath, "Create", err)
	}

	uid, err := res.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf(userPath, "Create", err)
	}

	return uid, nil
}
