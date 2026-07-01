5package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// CreateActivity registra una nueva sesión de entrenamiento para un usuario.
// TODO (Sprint 3): guardar la actividad, actualizar la racha y calcular el XP ganado.
func CreateActivity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "registro de actividad pendiente de implementación",
	})
}

// GetUserActivities devuelve el historial de actividades de un usuario.
// TODO (Sprint 3): obtener actividades desde la base de datos.
func GetUserActivities(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "historial de actividades pendiente de implementación",
		"user_id": id,
	})
}

// GetUserStreak devuelve la racha actual y el progreso de XP/nivel de un usuario.
// TODO (Sprint 3): calcular racha, XP y nivel desde la base de datos.
func GetUserStreak(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "racha de usuario pendiente de implementación",
		"user_id": id,
	})
}
