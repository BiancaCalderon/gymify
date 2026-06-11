package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthCheck responde con el estado del servicio. Útil para verificar
// que el backend está corriendo correctamente.
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "ok",
		"service": "gymify-backend",
	})
}
