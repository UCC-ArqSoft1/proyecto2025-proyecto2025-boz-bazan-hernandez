import api from './api';

export const enrollInActivity = async (activityId) => {
  try {
    const response = await api.post(`/enroll/${activityId}`);
    return response.data;
  } catch (error) {
    throw error.response?.data?.error || 'Error al inscribirse en la actividad';
  }
};

export const getUserActivities = async () => {
  try {
    const response = await api.get('/my-activities');
    return response.data;
  } catch (error) {
    throw error.response?.data?.error || 'Error al obtener tus actividades';
  }
};