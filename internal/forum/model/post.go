package model

import "time"

type Post struct {
	ID          uint64
	CommunityID uint64
	CreatedAt   time.Time
	UpgradedAt  time.Time

	Title      string
	Content    string
	VotesID    []uint64
	CommentsID []uint64

	Version uint64
}

// type PostCreateDTO struct {
// 	Title   string
// 	Content string
// }
