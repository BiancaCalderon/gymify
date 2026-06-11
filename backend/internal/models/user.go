package models

import "time"

// User representa a un usuario de Gymify.
type User struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Level        int       `json:"level"`
	XP           int       `json:"xp"`
	CreatedAt    time.Time `json:"created_at"`
}
