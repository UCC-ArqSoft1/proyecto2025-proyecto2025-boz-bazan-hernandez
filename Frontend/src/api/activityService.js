import api from './api';

export const getAllActivities = async () => {
  try {
    const response = await api.get('/activities');
    return response.data;
  } catch (error) {
    throw error.response?.data?.error || 'Error al obtener actividades';
  }
};

export const getActivityById = async (id) => {
  try {
    const response = await api.get(`/activities/${id}`);
    return response.data;
  } catch (error) {
    throw error.response?.data?.error || 'Error al obtener la actividad';
  }
};

export const createActivity = async (activityData) => {
  try {
    const response = await api.post('/activities', activityData);
    return response.data;
  } catch (error) {
    throw error.response?.data?.error || 'Error al crear la actividad';
  }
};

export const updateActivity = async (id, activityData) => {
  try {
    const response = await api.put(`/activities/${id}`, activityData);
    return response.data;
  } catch (error) {
    throw error.response?.data?.error || 'Error al actualizar la actividad';
  }
};

export const deleteActivity = async (id) => {
  try {
    const response = await api.delete(`/activities/${id}`);
    return response.data;
  } catch (error) {
    throw error.response?.data?.error || 'Error al eliminar la actividad';
  }
};

export const searchActivities = async (query) => {
  try {
    const response = await api.get('/activities/search', { params: { q: query } });
    return response.data;
  } catch (error) {
    throw error.response?.data?.error || 'Error al buscar actividades';
  }
};