package service

import (
	"errors"

	"forum/internal/forum/model"
	"forum/internal/forum/validator"
)

type UserStorage interface {
	Insert(user *model.CreateUserDTO) (uint64, error)
	GetUserByID(id uint64) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}

type User struct {
	storage UserStorage
}

func (u *User) Insert(user *model.CreateUserDTO) (uint64, error) {
	v := validator.New()
	if ValidateUser(v, user); !v.Valid() {
		return 0, errors.New("Invalid Credentials")
	}
	return u.storage.Insert(user)
}

func (u *User) GetUserByID(id uint64) (*model.User, error) {
	return u.storage.GetUserByID(id)
}

func (u *User) GetUserByEmail(email string) (*model.User, error) {
	v := validator.New()
	ValidateEmail(v, email)
	if !v.Valid() {
		return nil, errors.New("Invalid email")
	}
	return u.storage.GetUserByEmail(email)
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePassword(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 bytes long")
	v.Check(len(password) <= 72, "password", "must not be more than 72 bytes long")
}

func ValidateUser(v *validator.Validator, user *model.CreateUserDTO) {
	v.Check(user.Name != "", "name", "must be provided")
	v.Check(len(user.Name) <= 500, "name", "must not be more than 500 bytes long")

	ValidateEmail(v, user.Email)

	if user.Password != "" {
		ValidatePassword(v, user.Password)
	}
}
