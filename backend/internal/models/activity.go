package models

import "time"

// ActivityType representa el tipo de entrenamiento registrado.
type ActivityType string

const (
	ActivityWeights  ActivityType = "weights"
	ActivityCardio   ActivityType = "cardio"
	ActivityYoga     ActivityType = "yoga"
	ActivitySwimming ActivityType = "swimming"
	ActivityCycling  ActivityType = "cycling"
	ActivityOther    ActivityType = "other"
)

// Activity representa una sesión de entrenamiento registrada por un usuario.
type Activity struct {
	ID          string       `json:"id"`
	UserID      string       `json:"user_id"`
	Type        ActivityType `json:"type"`
	DurationMin int          `json:"duration_min"`
	Note        string       `json:"note,omitempty"`
	XPEarned    int          `json:"xp_earned"`
	CreatedAt   time.Time    `json:"created_at"`
}
