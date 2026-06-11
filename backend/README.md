# Gymify backend

API REST escrita en Go (Golang) para el sistema de gamificación social Gymify.

## Estructura

```
backend/
├── cmd/api/          Punto de entrada (main.go)
├── internal/
│   ├── config/       Carga de configuración / variables de entorno
│   ├── handlers/      Handlers HTTP (controladores)
│   ├── models/        Estructuras de datos del dominio
│   ├── repository/     Acceso a base de datos (pendiente)
│   └── middleware/      Middlewares personalizados (pendiente)
└── pkg/               Paquetes reutilizables
```

## Requisitos

- Go 1.22+
- PostgreSQL 15+

## Cómo correr el proyecto

1. Copia el archivo de variables de entorno:

   ```bash
   cp .env.example .env
   ```

2. Instala las dependencias:

   ```bash
   go mod tidy
   ```

3. Corre el servidor:

   ```bash
   go run ./cmd/api
   ```

4. Verifica que esté corriendo:

   ```bash
   curl http://localhost:8080/health
   ```

## Endpoints disponibles (Sprint 1)

Todos los endpoints están definidos pero devuelven `501 Not Implemented`
hasta que se conecten a la base de datos en sprints posteriores.

| Método | Ruta                              | Descripción                          | Sprint |
|--------|-----------------------------------|---------------------------------------|--------|
| GET    | `/health`                         | Estado del servicio                   | 1      |
| POST   | `/api/v1/auth/register`           | Registro de usuario                   | 2      |
| POST   | `/api/v1/auth/login`              | Inicio de sesión                      | 2      |
| GET    | `/api/v1/users/{id}`              | Obtener perfil de usuario             | 2      |
| POST   | `/api/v1/activities`              | Registrar actividad física            | 3      |
| GET    | `/api/v1/activities/user/{id}`    | Historial de actividades              | 3      |
| GET    | `/api/v1/streaks/user/{id}`       | Racha, XP y nivel del usuario         | 3      |
| GET    | `/api/v1/feed`                    | Feed social                           | 4      |
| POST   | `/api/v1/feed/posts`              | Crear publicación                     | 4      |
| GET    | `/api/v1/challenges`              | Listar retos grupales                 | 5      |
| POST   | `/api/v1/challenges/{id}/join`    | Unirse a un reto grupal               | 5      |

## Próximos pasos (Sprint 2)

- Configurar conexión a PostgreSQL (`internal/repository`)
- Implementar hashing de contraseñas y JWT
- Conectar `auth` y `users` a la base de datos
