package models

import "time"

// PostType representa el tipo de publicación en el feed social.
type PostType string

const (
	PostTraining   PostType = "training"
	PostAchievement PostType = "achievement"
)

// Post representa una publicación en el feed social de Gymify.
type Post struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Type      PostType  `json:"type"`
	Content   string    `json:"content"`
	LikeCount int       `json:"like_count"`
	CreatedAt time.Time `json:"created_at"`
}
