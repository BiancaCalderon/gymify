package models

import "time"

// Streak representa la racha de entrenamiento de un usuario.
// La racha se reinicia cada año para que los nuevos usuarios
// puedan integrarse a la experiencia sin quedar rezagados.
type Streak struct {
	UserID        string    `json:"user_id"`
	CurrentStreak int       `json:"current_streak"`
	LongestStreak int       `json:"longest_streak"`
	LastActiveDay time.Time `json:"last_active_day"`
	YearStarted   int       `json:"year_started"`
}
