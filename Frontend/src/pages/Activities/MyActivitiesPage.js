import React, { useState, useEffect } from 'react';
import ActivityList from '../../components/Activities/ActivityList';
import { getUserActivities } from '../../api/enrollmentService';
import Loading from '../../components/Common/Loading';

const MyActivitiesPage = () => {
  const [activities, setActivities] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchActivities = async () => {
      try {
        setLoading(true);
        const data = await getUserActivities();
        setActivities(data);
      } catch (error) {
        console.error('Error fetching user activities:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchActivities();
  }, []);

  if (loading) {
    return <Loading />;
  }

  return (
    <div className="animate__animated animate__fadeIn">
      <h1 className="text-3xl font-bold text-gray-800 mb-6">Mis Actividades</h1>
      <ActivityList activities={activities} />
    </div>
  );
};

export default MyActivitiesPage;