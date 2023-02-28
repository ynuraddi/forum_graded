package model

type Vote struct {
	ID     uint64
	UserID uint64
	PostID uint64
}

type VoteAddDTO struct {
	UserID uint64 `json:"user_id"`
	PostID uint64 `json:"post_id"`
}
