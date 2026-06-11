package models

import "time"

// Achievement representa un logro o insignia que un usuario puede desbloquear.
type Achievement struct {
	ID          string `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// UserAchievement representa un logro desbloqueado por un usuario específico.
type UserAchievement struct {
	UserID        string    `json:"user_id"`
	AchievementID string    `json:"achievement_id"`
	UnlockedAt    time.Time `json:"unlocked_at"`
}
