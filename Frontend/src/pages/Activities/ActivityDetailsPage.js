import React, { useState, useEffect } from 'react';
import { useParams, useHistory } from 'react-router-dom';
import ActivityDetail from '../../components/Activities/ActivityDetail';
import { getActivityById } from '../../api/activityService';
import { enrollInActivity } from '../../api/enrollmentService';
import Loading from '../../components/Common/Loading';
import { useContext } from 'react';
import { AuthContext } from '../../context/AuthContext';

const ActivityDetailsPage = () => {
  const { id } = useParams();
  const history = useHistory();
  const [activity, setActivity] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const { user } = useContext(AuthContext);

  useEffect(() => {
    const fetchActivity = async () => {
      try {
        setLoading(true);
        const data = await getActivityById(id);
        setActivity(data);
      } catch (error) {
        setError('Error al cargar la actividad');
        console.error('Error fetching activity:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchActivity();
  }, [id]);

  const handleEnroll = async () => {
    try {
      setError(null);
      await enrollInActivity(id);
      history.push('/my-activities');
    } catch (error) {
      setError(error);
    }
  };

  const clearError = () => {
    setError(null);
  };

  if (loading) {
    return <Loading />;
  }

  if (!activity) {
    return <div className="text-center py-10">Actividad no encontrada</div>;
  }

  return (
    <div className="animate__animated animate__fadeIn">
      <button
        onClick={() => history.goBack()}
        className="mb-4 text-indigo-600 hover:text-indigo-800 transition duration-300"
      >
        ← Volver
      </button>
      <ActivityDetail 
        activity={activity} 
        onEnroll={handleEnroll} 
        error={error} 
        onClearError={clearError} 
      />
    </div>
  );
};

export default ActivityDetailsPage;