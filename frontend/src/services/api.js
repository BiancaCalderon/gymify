import axios from 'axios';

// TODO: mover a variables de entorno (Sprint 2)
const API_BASE_URL = 'http://localhost:8080/api/v1';

export const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// TODO (Sprint 2): interceptor para adjuntar el token JWT a cada request
// api.interceptors.request.use((config) => {
//   const token = getStoredToken();
//   if (token) config.headers.Authorization = `Bearer ${token}`;
//   return config;
// });

export default api;
