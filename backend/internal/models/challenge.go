package models

import "time"

// ChallengeStatus representa el estado de un reto grupal.
type ChallengeStatus string

const (
	ChallengeUpcoming  ChallengeStatus = "upcoming"
	ChallengeActive    ChallengeStatus = "active"
	ChallengeCompleted ChallengeStatus = "completed"
)

// Challenge representa un reto grupal dentro de Gymify.
type Challenge struct {
	ID          string          `json:"id"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	XPReward    int             `json:"xp_reward"`
	Status      ChallengeStatus `json:"status"`
	StartDate   time.Time       `json:"start_date"`
	EndDate     time.Time       `json:"end_date"`
}

// ChallengeParticipant representa la participación de un usuario en un reto.
type ChallengeParticipant struct {
	ChallengeID string    `json:"challenge_id"`
	UserID      string    `json:"user_id"`
	Progress    int       `json:"progress"`
	JoinedAt    time.Time `json:"joined_at"`
}
