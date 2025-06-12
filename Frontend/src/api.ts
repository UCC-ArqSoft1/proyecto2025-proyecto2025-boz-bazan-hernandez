import axios, { AxiosInstance, AxiosResponse } from 'axios';
import {
    ActivityListItem,
    ActivityDetail,
    MyActivity,
    LoginRequest,
    LoginResponse,
    ApiErrorResponse,
    ApiSuccessResponse
} from './types';

class ApiService {
    private api: AxiosInstance;

    constructor() {
        this.api = axios.create({
            baseURL: process.env.REACT_APP_API_URL || 'http://localhost:8080/api',
            timeout: parseInt(process.env.REACT_APP_API_TIMEOUT || '10000'),
            headers: {
                'Content-Type': 'application/json',
            },
        });

        // Interceptor para agregar token automáticamente
        this.api.interceptors.request.use((config) => {
            const token = localStorage.getItem('token');
            if (token) {
                config.headers.Authorization = `Bearer ${token}`;
            }
            return config;
        });

        // Interceptor para manejar errores
        this.api.interceptors.response.use(
            (response) => response,
            (error) => {
                if (error.response?.status === 401) {
                    localStorage.removeItem('token');
                    localStorage.removeItem('user');
                }
                return Promise.reject(error);
            }
        );
    }


    async login(credentials: LoginRequest): Promise<LoginResponse> {
        const response: AxiosResponse<LoginResponse> = await this.api.post('/login', credentials);
        return response.data;
    }

    // GET /api/actividades (público)
    async getActivities(): Promise<ActivityListItem[]> {
        const response: AxiosResponse<ActivityListItem[]> = await this.api.get('/actividades');
        return response.data;
    }

    // GET /api/actividades/:id (público)
    async getActivityById(id: number): Promise<ActivityDetail> {
        const response: AxiosResponse<ActivityDetail> = await this.api.get(`/actividades/${id}`);
        return response.data;
    }

    // GET /api/mis-actividades (requiere auth)
    async getMyActivities(): Promise<MyActivity[]> {
        const response: AxiosResponse<MyActivity[]> = await this.api.get('/mis-actividades');
        return response.data;
    }

    // POST /api/inscribirse/:id (requiere auth)
    async enrollInActivity(activityId: number): Promise<ApiSuccessResponse> {
        const response: AxiosResponse<ApiSuccessResponse> = await this.api.post(`/inscribirse/${activityId}`);
        return response.data;
    }
}

export const apiService = new ApiService();