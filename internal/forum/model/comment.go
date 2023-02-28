package model

import "time"

type Comment struct {
	ID        uint64
	PostID    uint64
	AuthorID  uint64
	CreatedAt time.Time
}
