import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { ActivityDetail, ApiErrorResponse } from './types';
import { apiService } from './api';
import { isAuthenticated } from './utils';
import './ActivityDetail.css';

const ActivityDetailPage: React.FC = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [activity, setActivity] = useState<ActivityDetail | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string>('');
  const [enrolling, setEnrolling] = useState(false);
  const [enrollMessage, setEnrollMessage] = useState<{ type: 'success' | 'error', text: string } | null>(null);

  useEffect(() => {
    if (id) {
      loadActivity(parseInt(id));
    }
  }, [id]);

  const loadActivity = async (activityId: number) => {
    try {
      setLoading(true);
      setError('');
      const data = await apiService.getActivityById(activityId);
      setActivity(data);
    } catch (err: any) {
      console.error('Error loading activity:', err);
      setError('Error al cargar la actividad. Verifica que el backend esté ejecutándose.');
    } finally {
      setLoading(false);
    }
  };

  const handleEnroll = async () => {
    if (!activity) return;

    if (!isAuthenticated()) {
      setEnrollMessage({
        type: 'error',
        text: 'Debes iniciar sesión para inscribirte a una actividad'
      });
      return;
    }

    try {
      setEnrolling(true);
      setEnrollMessage(null);
      
      const response = await apiService.enrollInActivity(activity.id);
      
      setEnrollMessage({
        type: 'success',
        text: response.message // Backend envía {"message": "Inscripción exitosa"}
      });
      
      await loadActivity(activity.id);
      
    } catch (err: any) {
      console.error('Error enrolling:', err);
      
      let errorMessage = 'Error al inscribirse a la actividad';
      
      if (err.response?.data?.error) {
        errorMessage = err.response.data.error; // Backend envía {"error": "La actividad está llena"}
      } else if (err.message) {
        errorMessage = err.message;
      }
      
      setEnrollMessage({
        type: 'error',
        text: errorMessage
      });
    } finally {
      setEnrolling(false);
    }
  };

  const handleBackToHome = () => {
    navigate('/');
  };

  if (loading) {
    return (
      <div className="activity-detail-container">
        <div className="loading-spinner">
          <div className="spinner"></div>
          <p>Cargando actividad...</p>
        </div>
      </div>
    );
  }

  if (error || !activity) {
    return (
      <div className="activity-detail-container">
        <div className="error-message">
          <h2>⚠️ Error</h2>
          <p>{error || 'Actividad no encontrada'}</p>
          <button onClick={handleBackToHome} className="btn-back">
            Volver al inicio
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="activity-detail-container">
      <button onClick={handleBackToHome} className="btn-back-top">
        ← Volver a actividades
      </button>

      <div className="activity-detail-card">
        <div className="activity-image-section">
          <img 
            src={activity.foto_url || '/placeholder-activity.jpg'} 
            alt={activity.titulo}
            className="activity-image"
            onError={(e) => {
              (e.target as HTMLImageElement).src = '/placeholder-activity.jpg';
            }}
          />
        </div>

        <div className="activity-info-section">
          <header className="activity-header">
            <h1 className="activity-title">{activity.titulo}</h1>
            <div className="activity-badges">
              <span className="badge category-badge">{activity.categoria}</span>
              <span className="badge time-badge">{activity.horario}</span>
            </div>
          </header>

          <div className="activity-description">
            <p>{activity.descripcion}</p>
          </div>

          <div className="activity-details-grid">
            <div className="detail-item">
              <span className="detail-label">👨‍🏫 Instructor</span>
              <span className="detail-value">{activity.instructor}</span>
            </div>
            
            <div className="detail-item">
              <span className="detail-label">📅 Día</span>
              <span className="detail-value">{activity.dia}</span>
            </div>
            
            <div className="detail-item">
              <span className="detail-label">⏰ Horario</span>
              <span className="detail-value">{activity.horario}</span>
            </div>
            
            <div className="detail-item">
              <span className="detail-label">⏱️ Duración</span>
              <span className="detail-value">{activity.duracion}</span>
            </div>
            
            <div className="detail-item">
              <span className="detail-label">👥 Cupo Total</span>
              <span className="detail-value">{activity.cupo} personas</span>
            </div>
            
            <div className="detail-item">
              <span className="detail-label">🏷️ Categoría</span>
              <span className="detail-value">{activity.categoria}</span>
            </div>
          </div>

          {enrollMessage && (
            <div className={`alert ${enrollMessage.type === 'success' ? 'alert-success' : 'alert-error'}`}>
              {enrollMessage.text}
            </div>
          )}

          <div className="activity-actions">
            <button
              onClick={handleEnroll}
              disabled={enrolling}
              className="btn-enroll"
            >
              {enrolling ? 'Inscribiendo...' : 'Inscribirse a la actividad'}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};

export default ActivityDetailPage;
