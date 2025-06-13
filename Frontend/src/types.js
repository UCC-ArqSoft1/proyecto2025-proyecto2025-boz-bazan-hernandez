export const ActivityListItem = {
    id: Number,
    titulo: String,
    horario: String,
    profesor: String
};

export const ActivityDetail = {
    id: Number,
    titulo: String,
    descripcion: String,
    dia: String,
    horario: String,
    duracion: String,
    cupo: Number,
    categoria: String,
    instructor: String,
    foto_url: String
};

export const MyActivity = {
    id: Number,
    titulo: String,
    horario: String,
    dia: String,
    instructor: String
};

export const LoginRequest = {
    email: String,
    password: String
};

export const LoginResponse = {
    token: String,
    role: String
};

export const SearchFilters = {
    keyword: String,
    categoria: String,
    horario: String
};

export const ApiErrorResponse = {
    error: String
};

export const ApiSuccessResponse = {
    message: String
};
