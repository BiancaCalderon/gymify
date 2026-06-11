package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// GetUser devuelve el perfil de un usuario por su ID.
// TODO (Sprint 2): obtener el usuario desde la base de datos.
func GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "obtener usuario pendiente de implementación",
		"id":      id,
	})
}
