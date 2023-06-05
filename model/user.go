package model

type User struct {
	ID       int64
	Login    string
	Email    string
	Password string
	IsActive bool
	Version  string
}
