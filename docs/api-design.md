# Diseño Inicial de API - Gymify

## Autenticación

### Registro

POST /auth/register

Body:

{
  "name": "Bianca",
  "email": "bianca@email.com",
  "password": "password123"
}

---

### Inicio de Sesión

POST /auth/login

Body:

{
  "email": "bianca@email.com",
  "password": "password123"
}

---

## Usuarios

### Obtener Perfil

GET /users/profile

---

### Actualizar Perfil

PUT /users/profile

---

## Actividades

### Registrar Entrenamiento

POST /activities

Body:

{
  "activity_type": "Gym",
  "duration": 60
}

---

### Obtener Historial

GET /activities

---

## Logros

### Obtener Logros

GET /achievements

---

### Obtener Logros del Usuario

GET /users/achievements

---

## Publicaciones

### Crear Publicación

POST /posts

Body:

{
  "caption": "Nuevo récord en sentadilla",
  "image_url": "url-imagen"
}

---

### Obtener Feed

GET /posts

---

### Obtener Publicación

GET /posts/{id}

---

## Comentarios

### Crear Comentario

POST /posts/{id}/comments

---

## Likes

### Dar Like

POST /posts/{id}/like

---

### Quitar Like

DELETE /posts/{id}/like

---

## Retos

### Obtener Retos

GET /challenges

---

### Unirse a un Reto

POST /challenges/{id}/join

---

### Ver Participantes

GET /challenges/{id}/participants

---

## Rankings

### Ranking General

GET /rankings

---

## Moderación

### Reportar Publicación

POST /posts/{id}/report

Body:

{
  "reason": "Contenido inapropiado"
}

---

## Notificaciones

### Obtener Notificaciones

GET /notifications