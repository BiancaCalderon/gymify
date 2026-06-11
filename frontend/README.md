# Gymify frontend

App móvil híbrida de Gymify, construida con React Native y Expo.

## Estructura

```
frontend/
├── App.js                  Punto de entrada
├── src/
│   ├── screens/            Pantallas de la app
│   ├── components/         Componentes reutilizables
│   ├── navigation/          Navegación (stack + tabs)
│   ├── services/            Cliente API (axios)
│   ├── constants/           Tema (colores, espaciados, tipografía)
│   └── assets/              Íconos, imágenes
```

## Requisitos

- Node.js 18+
- Expo CLI (`npm install -g expo-cli`) o usar `npx expo`
- App Expo Go en tu teléfono (para probar en dispositivo físico)

## Cómo correr el proyecto

1. Instala las dependencias:

   ```bash
   npm install
   ```

2. Inicia el servidor de desarrollo:

   ```bash
   npm start
   ```

3. Escanea el código QR con la app Expo Go (Android) o la cámara (iOS).

## Pantallas incluidas (Sprint 1)

| Pantalla              | Estado                          | Sprint de implementación |
|-----------------------|----------------------------------|---------------------------|
| Login                 | UI completa, falta conectar API  | 2                          |
| Registro              | Placeholder                       | 2                          |
| Home / Dashboard      | UI completa con datos de ejemplo | 3                          |
| Registrar actividad   | Placeholder                       | 3                          |
| Feed social           | Placeholder                       | 4                          |
| Retos grupales        | Placeholder                       | 5                          |
| Perfil                | Placeholder                       | 3-4                        |

## Sistema de diseño

Definido en `src/constants/theme.js`. Dark mode con acento naranja (`#FF6B35`)
y dorado para XP (`#FFB347`).

## Próximos pasos (Sprint 2)

- Conectar `LoginScreen` y `RegisterScreen` a `POST /api/v1/auth/login` y `/register`
- Guardar el token JWT (e.g. con `expo-secure-store`)
- Mostrar `AuthStack` o `MainTabs` según el estado de sesión
