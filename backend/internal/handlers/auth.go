package handlers

import (
	"encoding/json"
	"net/http"
)

// RegisterRequest representa el cuerpo esperado para crear una cuenta nueva.
type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginRequest representa el cuerpo esperado para iniciar sesión.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Register crea un nuevo usuario.
// TODO (Sprint 2): validar datos, hashear contraseña y guardar en la base de datos.
func Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "cuerpo de la petición inválido", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "registro pendiente de implementación",
	})
}

// Login autentica a un usuario y devuelve un token JWT.
// TODO (Sprint 2): validar credenciales contra la base de datos y firmar el JWT.
func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "cuerpo de la petición inválido", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "login pendiente de implementación",
	})
}
