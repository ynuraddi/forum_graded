package service

import "forum/internal/forum/model"

type UserStorage interface {
	Insert(user *model.User) error
	GetUserByID(id uint64) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

type User struct {
	storage UserStorage
}

func Insert(user *model.User) error {
	return nil
}

func GetUserByID(id uint64) (*model.User, error) {
	return nil, nil
}

func GetUserByEmail(email string) (*model.User, error) {
	return nil, nil
}
