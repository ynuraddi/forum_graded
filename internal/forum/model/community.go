package model

type Community struct {
	ID uint64
	// Type        bool // 0 - private, 1 - public
	Subscribers []uint64

	Title      string
	Desciption string

	PostID []uint64
}
