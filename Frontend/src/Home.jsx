import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { apiService } from './api';
import ActivityCard from './ActivityCard';
import SearchBar from './SearchBar';
import './ActivityCard.css';
import './SearchBar.css';
import './Home.css';

const Home = () => {
  const [activities, setActivities] = useState([]);
  const [filteredActivities, setFilteredActivities] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    loadActivities();
  }, []);

  const loadActivities = async () => {
    try {
      setLoading(true);
      setError('');
      const data = await apiService.getActivities();
      setActivities(data);
      setFilteredActivities(data);
    } catch (err) {
      console.error('Error loading activities:', err);
      setError('Error al cargar las actividades. Verifica que el backend esté ejecutándose en puerto 8080.');
    } finally {
      setLoading(false);
    }
  };

  const handleSearch = (filters) => {
    let filtered = [...activities];

    if (filters.keyword && filters.keyword.trim()) {
      const keyword = filters.keyword.toLowerCase();
      filtered = filtered.filter(activity => 
        activity.titulo.toLowerCase().includes(keyword) ||
        activity.profesor.toLowerCase().includes(keyword)
      );
    }

    if (filters.categoria && filters.categoria.trim()) {
      const categoria = filters.categoria.toLowerCase();
      filtered = filtered.filter(activity => 
        activity.titulo.toLowerCase().includes(categoria)
      );
    }

    if (filters.horario && filters.horario.trim()) {
      const horario = filters.horario;
      filtered = filtered.filter(activity => {
        const horaActividad = parseInt(activity.horario.split(':')[0]);
        
        if (horario.includes('Mañana')) {
          return horaActividad >= 6 && horaActividad < 12;
        } else if (horario.includes('Tarde')) {
          return horaActividad >= 12 && horaActividad < 18;
        } else if (horario.includes('Noche')) {
          return horaActividad >= 18 && horaActividad <= 22;
        }
        return true;
      });
    }

    setFilteredActivities(filtered);
  };

  const handleClearSearch = () => {
    setFilteredActivities(activities);
  };

  const handleActivityClick = (id) => {
    navigate(`/actividad/${id}`);
  };

  if (loading) {
    return (
      <div className="home-container">
        <div className="loading-spinner">
          <div className="spinner"></div>
          <p>Cargando actividades...</p>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="home-container">
        <div className="error-message">
          <h2>⚠️ Error de conexión</h2>
          <p>{error}</p>
          <button onClick={loadActivities} className="btn-retry">
            Intentar nuevamente
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="home-container">
      <header className="home-header">
        <h1 className="home-title">Actividades Deportivas</h1>
        <p className="home-subtitle">
          Descubre y únete a nuestras actividades deportivas
        </p>
      </header>

      <SearchBar onSearch={handleSearch} onClear={handleClearSearch} />

      <div className="activities-section">
        {filteredActivities.length === 0 ? (
          <div className="no-results">
            <h3>No se encontraron actividades</h3>
            <p>Intenta ajustar los filtros de búsqueda</p>
          </div>
        ) : (
          <>
            <div className="results-header">
              <p className="results-count">
                {filteredActivities.length} actividad{filteredActivities.length !== 1 ? 'es' : ''} encontrada{filteredActivities.length !== 1 ? 's' : ''}
              </p>
            </div>
            
            <div className="activities-grid">
              {filteredActivities.map((activity) => (
                <ActivityCard
                  key={activity.id}
                  activity={activity}
                  onClick={handleActivityClick}
                />
              ))}
            </div>
          </>
        )}
      </div>
    </div>
  );
};

export default Home;
