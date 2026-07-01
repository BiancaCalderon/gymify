package config

import "os"

// Config contiene la configuración de la aplicación cargada desde variables de entorno.
type Config struct {
	Env         string
	DatabaseURL string
	JWTSecret   string
}

// Load lee las variables de entorno y devuelve la configuración de la app.
// Si una variable no está definida, se usa un valor por defecto para desarrollo local.
func Load() *Config {
	return &Config{
		Env:         getEnv("APP_ENV", "development"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://gymify:gymify@localhost:5432/gymify?sslmode=disable"),
		JWTSecret:   getEnv("JWT_SECRET", "change-me-in-production"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
