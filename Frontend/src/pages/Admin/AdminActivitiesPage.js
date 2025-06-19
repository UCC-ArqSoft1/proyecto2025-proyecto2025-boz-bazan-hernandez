import React, { useState, useEffect } from 'react';
import ActivityList from '../../components/Activities/ActivityList';
import { getAllActivities, deleteActivity } from '../../api/activityService';
import Loading from '../../components/Common/Loading';
import { useHistory } from 'react-router-dom';
import Alert from '../../components/Common/Alert';

const AdminActivitiesPage = () => {
  const [activities, setActivities] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const history = useHistory();

  useEffect(() => {
    const fetchActivities = async () => {
      try {
        setLoading(true);
        const data = await getAllActivities();
        setActivities(data);
      } catch (error) {
        console.error('Error fetching activities:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchActivities();
  }, []);

  const handleEdit = (id) => {
    history.push(`/admin/activities/edit/${id}`);
  };

  const handleDelete = async (id) => {
    try {
      setError(null);
      await deleteActivity(id);
      setActivities(activities.filter(activity => activity.id !== id));
    } catch (error) {
      setError('Error al eliminar la actividad');
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
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold text-gray-800">Administrar Actividades</h1>
        <button
          onClick={() => history.push('/admin/activities/create')}
          className="bg-green-600 text-white px-4 py-2 rounded-md hover:bg-green-700 transition duration-300"
        >
          Crear Nueva Actividad
        </button>
      </div>
      
      {error && <Alert message={error} type="error" onClose={clearError} />}
      
      <div className="bg-white rounded-lg shadow-md overflow-hidden">
        <table className="min-w-full divide-y divide-gray-200">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Título</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Horario</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Instructor</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Acciones</th>
            </tr>
          </thead>
          <tbody className="bg-white divide-y divide-gray-200">
            {activities.map((activity) => (
              <tr key={activity.id}>
                <td className="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900">{activity.titulo}</td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{activity.horario}</td>
                <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">{activity.instructor}</td>
                <td className="px-6 py-4 whitespace-nowrap text-sm font-medium">
                  <button
                    onClick={() => handleEdit(activity.id)}
                    className="text-indigo-600 hover:text-indigo-900 mr-4"
                  >
                    Editar
                  </button>
                  <button
                    onClick={() => handleDelete(activity.id)}
                    className="text-red-600 hover:text-red-900"
                  >
                    Eliminar
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default AdminActivitiesPage;