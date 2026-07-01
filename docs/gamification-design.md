# Gymify

Sistema de gamificación social para mejorar la experiencia del entrenamiento en usuarios de gimnasios en Guatemala.

## Descripción

Gymify es una aplicación móvil híbrida desarrollada como proyecto de tesis para la Universidad del Valle de Guatemala.

El objetivo del proyecto es incrementar la motivación y la adherencia al entrenamiento mediante la incorporación de elementos de gamificación social como:

- Rachas de entrenamiento (Streaks)
- Sistema de experiencia (XP)
- Niveles
- Logros (Achievements)
- Feed social
- Retos grupales
- Rankings
- Reacciones y comentarios
- Notificaciones

La aplicación busca combinar principios utilizados por plataformas como Duolingo e Instagram adaptados al contexto fitness.

---

# Objetivo General

Implementar un sistema de gamificación social en una aplicación móvil híbrida que incremente la adherencia y motivación en usuarios frecuentes de gimnasio en Guatemala.

---

# Tecnologías

## Frontend

- React Native
- Expo
- TypeScript
- React Navigation
- Zustand (estado global)
- Axios

## Backend

- Golang
- Gin Framework
- GORM
- JWT Authentication

## Base de datos

- PostgreSQL

---

# Arquitectura

El proyecto sigue una arquitectura cliente-servidor.

```
React Native
        │
 REST API
        │
   Backend Go
        │
 PostgreSQL
```

---

# Estructura del proyecto

```
gymify/

├── backend/
│   ├── cmd/
│   ├── internal/
│   ├── pkg/
│   ├── configs/
│   ├── migrations/
│   └── docs/
│
├── frontend/
│   ├── src/
│   ├── assets/
│   └── docs/
│
├── database/
│
├── docs/
│
└── README.md
```

---

# Funcionalidades

## Gestión de usuarios

- Registro
- Inicio de sesión
- Perfil
- Configuración

## Entrenamiento

- Registrar entrenamiento
- Historial
- Calendario
- Estadísticas

## Gamificación

- XP
- Niveles
- Streaks
- Logros
- Rankings

## Comunidad

- Feed social
- Compartir logros
- Compartir fotografías
- Comentarios
- Reacciones
- Seguir usuarios

## Retos

- Retos individuales
- Retos grupales
- Clasificaciones

---

# Sprint actual

Sprint 1

- Levantamiento de requerimientos
- Historias de usuario
- Diseño UX/UI
- Arquitectura
- Diseño de base de datos
- Configuración del entorno

---

# Estado del proyecto

 En desarrollo


