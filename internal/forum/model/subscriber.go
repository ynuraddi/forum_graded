package model

type Subscriber struct {
	UserID      uint64
	CommunityID uint64
	Permission  uint8 // 0 - admin, 1 - moderator, 2 - subscriber
}
