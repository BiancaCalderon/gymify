# Diseño Preliminar de Base de Datos - Gymify

## Entidades Principales

### Users

Almacena la información general de cada usuario.

| Campo | Tipo |
|---------|---------|
| id | UUID |
| name | VARCHAR |
| email | VARCHAR |
| password_hash | VARCHAR |
| level | INTEGER |
| xp | INTEGER |
| streak | INTEGER |
| created_at | TIMESTAMP |

---

### Activities

Registra los entrenamientos realizados por los usuarios.

| Campo | Tipo |
|---------|---------|
| id | UUID |
| user_id | UUID |
| activity_type | VARCHAR |
| duration_minutes | INTEGER |
| created_at | TIMESTAMP |

---

### Achievements

Define los logros disponibles dentro del sistema.

| Campo | Tipo |
|---------|---------|
| id | UUID |
| title | VARCHAR |
| description | TEXT |
| xp_reward | INTEGER |

---

### UserAchievements

Relación entre usuarios y logros obtenidos.

| Campo | Tipo |
|---------|---------|
| user_id | UUID |
| achievement_id | UUID |
| unlocked_at | TIMESTAMP |

---

### Posts

Publicaciones realizadas por los usuarios.

| Campo | Tipo |
|---------|---------|
| id | UUID |
| user_id | UUID |
| image_url | VARCHAR |
| caption | TEXT |
| created_at | TIMESTAMP |

---

### Comments

Comentarios en publicaciones.

| Campo | Tipo |
|---------|---------|
| id | UUID |
| post_id | UUID |
| user_id | UUID |
| content | TEXT |
| created_at | TIMESTAMP |

---

### Likes

Reacciones a publicaciones.

| Campo | Tipo |
|---------|---------|
| id | UUID |
| post_id | UUID |
| user_id | UUID |

---

### Challenges

Retos grupales disponibles.

| Campo | Tipo |
|---------|---------|
| id | UUID |
| title | VARCHAR |
| description | TEXT |
| xp_reward | INTEGER |
| start_date | DATE |
| end_date | DATE |

---

### ChallengeParticipants

Usuarios inscritos en retos.

| Campo | Tipo |
|---------|---------|
| challenge_id | UUID |
| user_id | UUID |
| joined_at | TIMESTAMP |

---

## Relaciones

- Un usuario puede registrar múltiples actividades.
- Un usuario puede obtener múltiples logros.
- Un usuario puede crear múltiples publicaciones.
- Una publicación puede tener múltiples comentarios.
- Una publicación puede recibir múltiples likes.
- Un usuario puede participar en múltiples retos.
- Un reto puede tener múltiples participantes.