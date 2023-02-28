package model

type User struct {
	ID           uint64
	Name         string
	HashPassword string
	Email        string
}
