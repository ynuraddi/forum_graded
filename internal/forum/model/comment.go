package model

import "time"

type Comment struct {
	ID        uint64
	PostID    uint64
	AuthorID  uint64
	CreatedAt time.Time
}

type CommentAddDTO struct {
	PostID   uint64 `json:"post_id"`
	AuthorID uint64 `json:"author_id"`
}
