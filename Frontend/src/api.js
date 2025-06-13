import axios from 'axios';

class ApiService {
    constructor() {
        this.api = axios.create({
            baseURL: process.env.REACT_APP_API_URL || 'http://localhost:8080/api',
            timeout: parseInt(process.env.REACT_APP_API_TIMEOUT || '10000'),
            headers: {
                'Content-Type': 'application/json',
            },
        });

        this.api.interceptors.request.use((config) => {
            const token = localStorage.getItem('token');
            if (token) {
                config.headers.Authorization = `Bearer ${token}`;
            }
            return config;
        });

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

    async login(credentials) {
        const response = await this.api.post('/login', credentials);
        return response.data;
    }

    async getActivities() {
        const response = await this.api.get('/actividades');
        return response.data;
    }

    async getActivityById(id) {
        const response = await this.api.get(`/actividades/${id}`);
        return response.data;
    }

    async getMyActivities() {
        const response = await this.api.get('/mis-actividades');
        return response.data;
    }

    async enrollInActivity(activityId) {
        const response = await this.api.post(`/inscribirse/${activityId}`);
        return response.data;
    }
}

export const apiService = new ApiService();
