import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import ActivityForm from '../../components/Activities/ActivityForm';
import { createActivity } from '../../api/activityService';
import Loading from '../../components/Common/Loading';
import Alert from '../../components/Common/Alert';

const CreateActivityPage = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const history = useHistory();

  const handleSubmit = async (activityData) => {
    try {
      setLoading(true);
      setError(null);
      await createActivity(activityData);
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

  return (
    <div className="animate__animated animate__fadeIn">
      <h1 className="text-3xl font-bold text-gray-800 mb-6">Crear Nueva Actividad</h1>
      <ActivityForm onSubmit={handleSubmit} error={error} onClearError={clearError} />
    </div>
  );
};

export default CreateActivityPage;