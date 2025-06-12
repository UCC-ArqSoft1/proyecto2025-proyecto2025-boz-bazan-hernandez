export interface ActivityListItem {
    id: number;
    titulo: string;
    horario: string;
    profesor: string; // ← Backend mapea instructor → profesor
}

// Para GET /api/actividades/:id (ActivityDetailResponse del backend)
export interface ActivityDetail {
    id: number;
    titulo: string;
    descripcion: string;
    dia: string;
    horario: string;
    duracion: string;
    cupo: number;
    categoria: string;
    instructor: string;
    foto_url: string;
}

// Para GET /api/mis-actividades (MyActivityResponse del backend)
export interface MyActivity {
    id: number;
    titulo: string;
    horario: string;
    dia: string;
    instructor: string;
}

// Para POST /api/login
export interface LoginRequest {
    email: string;
    password: string;
}

export interface LoginResponse {
    token: string;
    role: string; // "socio" o "administrador"
}

// Para búsqueda (frontend only)
export interface SearchFilters {
    keyword?: string;
    categoria?: string;
    horario?: string;
}

// Para respuestas de error del backend
export interface ApiErrorResponse {
    error: string;
}

// Para respuestas de éxito del backend
export interface ApiSuccessResponse {
    message: string;
}