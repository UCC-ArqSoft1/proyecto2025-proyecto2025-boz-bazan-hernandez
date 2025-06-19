// Home.js - Versión que depende únicamente del backend
import React, { useState, useEffect } from 'react';
import ActivityList from '../../components/Activities/ActivityList';
import SearchBar from '../../components/Common/SearchBar';
import { getAllActivities, searchActivities } from '../../api/activityService';
import Loading from '../../components/Common/Loading';
import './Home.css'; // Importar los estilos del Home

const Home = () => {
  const [activities, setActivities] = useState([]);
  const [loading, setLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');
  const [filteredActivities, setFilteredActivities] = useState([]);
  const [selectedCategory, setSelectedCategory] = useState('all');
  const [error, setError] = useState(null);

  // Categorías disponibles - se calculan dinámicamente desde el backend
  const [categories, setCategories] = useState([
    { id: 'all', name: 'Todas', icon: '🌟', count: 0 }
  ]);

  useEffect(() => {
    const fetchActivities = async () => {
      try {
        setLoading(true);
        setError(null);
        let data;
        
        if (searchQuery) {
          data = await searchActivities(searchQuery);
        } else {
          data = await getAllActivities();
        }
        
        setActivities(data);
        setFilteredActivities(data);
        
        // Calcular categorías dinámicamente desde los datos del backend
        if (data && data.length > 0) {
          const uniqueCategories = [...new Set(data.map(activity => activity.category))];
          const dynamicCategories = [
            { id: 'all', name: 'Todas', icon: '🌟', count: data.length }
          ];
          
          uniqueCategories.forEach(category => {
            const count = data.filter(activity => activity.category === category).length;
            const categoryInfo = getCategoryInfo(category);
            dynamicCategories.push({
              id: category,
              name: categoryInfo.name,
              icon: categoryInfo.icon,
              count: count
            });
          });
          
          setCategories(dynamicCategories);
        } else {
          // Si no hay actividades, resetear categorías
          setCategories([{ id: 'all', name: 'Todas', icon: '🌟', count: 0 }]);
        }
        
      } catch (error) {
        console.error('Error fetching activities:', error);
        setError('Error al cargar las actividades. Por favor, intenta nuevamente.');
        setActivities([]);
        setFilteredActivities([]);
        setCategories([{ id: 'all', name: 'Todas', icon: '🌟', count: 0 }]);
      } finally {
        setLoading(false);
      }
    };

    fetchActivities();
  }, [searchQuery]);

  // Función para obtener información de categoría basada en el nombre
  const getCategoryInfo = (categoryName) => {
    const categoryMap = {
      'fitness': { name: 'Fitness', icon: '💪' },
      'yoga': { name: 'Yoga', icon: '🧘‍♀️' },
      'natacion': { name: 'Natación', icon: '🏊‍♂️' },
      'baile': { name: 'Baile', icon: '💃' },
      'deportes': { name: 'Deportes', icon: '⚽' },
      'pilates': { name: 'Pilates', icon: '🤸‍♀️' },
      'running': { name: 'Running', icon: '🏃‍♂️' },
      'crossfit': { name: 'CrossFit', icon: '🏋️‍♂️' },
      'boxeo': { name: 'Boxeo', icon: '🥊' },
      'tenis': { name: 'Tenis', icon: '🎾' },
      'basquet': { name: 'Básquet', icon: '🏀' },
      'futbol': { name: 'Fútbol', icon: '⚽' },
      'spinning': { name: 'Spinning', icon: '🚴‍♂️' },
      'zumba': { name: 'Zumba', icon: '💃' },
      'aqua': { name: 'Aqua Aeróbicos', icon: '🌊' }
    };
    
    return categoryMap[categoryName.toLowerCase()] || { name: categoryName, icon: '🏃‍♂️' };
  };

  // Filtrar por categoría
  useEffect(() => {
    if (selectedCategory === 'all') {
      setFilteredActivities(activities);
    } else {
      const filtered = activities.filter(activity => 
        activity.category === selectedCategory
      );
      setFilteredActivities(filtered);
    }
  }, [activities, selectedCategory]);

  const handleSearch = (query) => {
    setSearchQuery(query);
    setSelectedCategory('all');
  };

  const handleCategoryChange = (categoryId) => {
    setSelectedCategory(categoryId);
    setSearchQuery('');
  };

  const handleRetry = () => {
    setError(null);
    setSearchQuery('');
    setSelectedCategory('all');
    // El useEffect se ejecutará automáticamente al cambiar searchQuery
  };

  // Calcular estadísticas reales desde el backend
  const stats = {
    totalActivities: activities.length,
    totalCategories: categories.length - 1, // -1 para excluir "Todas"
    averagePrice: activities.length > 0 
      ? Math.round(activities.reduce((sum, a) => sum + (a.price || 0), 0) / activities.length)
      : 0,
    totalInstructors: activities.length > 0 
      ? [...new Set(activities.map(a => a.instructor).filter(Boolean))].length
      : 0
  };

  if (loading) {
    return (
      <div className="home-container">
        <div className="home-loading">
          <div className="home-loading-spinner"></div>
          <div className="home-loading-text">Cargando actividades desde el servidor...</div>
          <div style={{ marginTop: '1rem', fontSize: '0.9rem', opacity: 0.8 }}>
            Conectando con el backend ⚡
          </div>
        </div>
      </div>
    );
  }

  if (error) {
    return (
      <div className="home-container">
        <div className="home-content">
          <div className="error-state">
            <div className="error-icon">⚠️</div>
            <h2 className="error-title">Error de Conexión</h2>
            <p className="error-message">{error}</p>
            <button onClick={handleRetry} className="btn btn-primary retry-button">
              <span className="button-icon">🔄</span>
              Reintentar
            </button>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="home-container">
      <div className="home-content">
        {/* Título principal */}
        <h1 className="home-title">
          Descubre Actividades Increíbles
        </h1>
        <p className="home-subtitle">
          Encuentra la actividad perfecta para ti. Desde deportes hasta talleres creativos, 
          tenemos algo especial esperándote. ¡Comienza tu aventura fitness hoy!
        </p>

        {/* Estadísticas reales del backend */}
        {activities.length > 0 && (
          <div className="stats-section">
            <div className="stat-card">
              <span className="stat-number">{stats.totalActivities}</span>
              <span className="stat-label">Actividades</span>
            </div>
            <div className="stat-card">
              <span className="stat-number">{stats.totalCategories}</span>
              <span className="stat-label">Categorías</span>
            </div>
            <div className="stat-card">
              <span className="stat-number">{stats.totalInstructors}</span>
              <span className="stat-label">Instructores</span>
            </div>
            <div className="stat-card">
              <span className="stat-number">${stats.averagePrice}</span>
              <span className="stat-label">Precio Promedio</span>
            </div>
          </div>
        )}

        {/* Sección de búsqueda */}
        <div className="search-section">
          <h2>🔍 Busca tu actividad ideal</h2>
          <SearchBar onSearch={handleSearch} />
        </div>

        {/* Filtros de categoría - solo si hay actividades */}
        {activities.length > 0 && categories.length > 1 && (
          <div className="category-filters">
            <h3 className="category-title">Explora por categorías</h3>
            <div className="category-grid">
              {categories.map(category => (
                <button
                  key={category.id}
                  className={`category-card ${selectedCategory === category.id ? 'active' : ''}`}
                  onClick={() => handleCategoryChange(category.id)}
                >
                  <div className="category-icon">{category.icon}</div>
                  <div className="category-name">{category.name}</div>
                  <div className="category-count">{category.count} actividad{category.count !== 1 ? 'es' : ''}</div>
                </button>
              ))}
            </div>
          </div>
        )}

        {/* Sección de actividades */}
        <div className="activities-section">
          <div className="activities-header">
            <h2 className="activities-title">
              {searchQuery 
                ? `Resultados para "${searchQuery}"` 
                : selectedCategory === 'all' 
                  ? 'Todas las Actividades'
                  : `Actividades de ${categories.find(c => c.id === selectedCategory)?.name}`
              }
            </h2>
            <p className="activities-subtitle">
              {filteredActivities.length === 0 
                ? 'No se encontraron actividades' 
                : `${filteredActivities.length} actividad${filteredActivities.length !== 1 ? 'es' : ''} disponible${filteredActivities.length !== 1 ? 's' : ''}`
              }
            </p>
          </div>
          
          {filteredActivities.length === 0 ? (
            <div className="empty-state">
              <div className="empty-state-icon">🏃‍♂️</div>
              <h3 className="empty-state-title">
                {activities.length === 0 
                  ? 'No hay actividades disponibles en el servidor'
                  : 'No se encontraron actividades con estos criterios'
                }
              </h3>
              <p className="empty-state-subtitle">
                {activities.length === 0 
                  ? 'El backend no ha devuelto ninguna actividad. Verifica la conexión con el servidor.'
                  : searchQuery 
                    ? 'Intenta con otros términos de búsqueda o explora todas las actividades.'
                    : 'No hay actividades en esta categoría.'
                }
              </p>
              {(searchQuery || selectedCategory !== 'all') && activities.length > 0 && (
                <button 
                  className="btn btn-primary"
                  onClick={() => {
                    setSearchQuery('');
                    setSelectedCategory('all');
                  }}
                  style={{ marginTop: '1rem' }}
                >
                  Ver todas las actividades
                </button>
              )}
            </div>
          ) : (
            <ActivityList activities={filteredActivities} />
          )}
        </div>

        {/* Call to action - solo si hay actividades */}
        {activities.length > 0 && (
          <div className="cta-section">
            <div className="cta-content">
              <h2 className="cta-title">¿Listo para comenzar tu aventura fitness?</h2>
              <p className="cta-subtitle">
                Únete a miles de personas que ya están transformando sus vidas con nuestras actividades.
              </p>
              <div className="cta-buttons">
                <button 
                  className="btn btn-primary cta-button"
                  onClick={() => setSelectedCategory('all')}
                >
                  <span className="button-icon">🎯</span>
                  Explorar Todas las Actividades
                </button>
              </div>
            </div>
            <div className="cta-image">
              <div className="cta-emoji">🏆</div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default Home;