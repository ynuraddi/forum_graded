package repository

import "forum/internal/forum/model"

type Users interface {
	Insert(user *model.User) error
	GetUserByID(id uint64) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}
