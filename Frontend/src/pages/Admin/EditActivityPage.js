
import React, { useState, useEffect } from 'react';
import { useParams, useHistory } from 'react-router-dom';
import ActivityForm from '../../components/Activities/ActivityForm';
import { getActivityById, updateActivity } from '../../api/activityService';
import Loading from '../../components/Common/Loading';
import Alert from '../../components/Common/Alert';

const EditActivityPage = () => {
  const { id } = useParams();
  const history = useHistory();
  const [activity, setActivity] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

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

  const handleSubmit = async (activityData) => {
    try {
      setLoading(true);
      setError(null);
      await updateActivity(id, activityData);
      history.push('/admin/activities');
    } catch (error) {
      setError(error);
    } finally {
      setLoading(false);
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

  // Convert activity to form data format
  const initialFormData = {
    titulo: activity.titulo,
    descripcion: activity.descripcion,
    categoria: activity.categoria,
    instructor: activity.instructor,
    dia: activity.dia,
    horario: activity.horario,
    duracion: `${activity.duracion} min`,
    cupo: activity.cupo,
    fotoURL: activity.fotoURL,
  };

  return (
    <div className="animate__animated animate__fadeIn">
      <h1 className="text-3xl font-bold text-gray-800 mb-6">Editar Actividad</h1>
      <ActivityForm 
        initialData={initialFormData} 
        onSubmit={handleSubmit} 
        error={error} 
        onClearError={clearError} 
      />
    </div>
  );
};

export default EditActivityPage;