package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// GetChallenges devuelve la lista de retos grupales disponibles y activos.
// TODO (Sprint 5): obtener retos desde la base de datos, filtrando por estado.
func GetChallenges(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "lista de retos pendiente de implementación",
	})
}

// JoinChallenge inscribe a un usuario en un reto grupal.
// TODO (Sprint 5): registrar la participación del usuario en el reto.
func JoinChallenge(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"message":      "unirse a reto pendiente de implementación",
		"challenge_id": id,
	})
}
