package handlers

import (
	"encoding/json"
	"net/http"
)

// GetFeed devuelve las publicaciones recientes de la comunidad.
// TODO (Sprint 4): obtener posts de usuarios seguidos desde la base de datos.
func GetFeed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "feed social pendiente de implementación",
	})
}

// CreatePost crea una nueva publicación (logro, entrenamiento, etc.).
// TODO (Sprint 4): validar contenido y guardar el post en la base de datos.
func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "creación de post pendiente de implementación",
	})
}
