package models

import "time"

// Comment structure
type Comment struct {
	ID        int64      `json:"id"`
	Content   string     `json:"content"`
	UserID    int64      `json:"user_id"`
	Username  string     `json:"username"`
	AvatarURL string     `json:"avatar_url"`
	PostID    int64      `json:"post_id"`
	CreatedAt time.Time  `json:"created_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

// ClientStream structure
type ClientStream struct {
	UserID  int64
	Comment chan *Comment
}
