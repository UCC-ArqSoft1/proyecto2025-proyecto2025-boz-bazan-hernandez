import React, { useState, useEffect } from 'react';
import ActivityList from '../../components/Activities/ActivityList';
import SearchBar from '../../components/Common/SearchBar';
import { getAllActivities, searchActivities } from '../../api/activityService';
import Loading from '../../components/Common/Loading';

const Activities = () => {
  const [activities, setActivities] = useState([]);
  const [loading, setLoading] = useState(true);
  const [searchQuery, setSearchQuery] = useState('');

  useEffect(() => {
    const fetchActivities = async () => {
      try {
        setLoading(true);
        let data;
        if (searchQuery) {
          data = await searchActivities(searchQuery);
        } else {
          data = await getAllActivities();
        }
        setActivities(data);
      } catch (error) {
        console.error('Error fetching activities:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchActivities();
  }, [searchQuery]);

  const handleSearch = (query) => {
    setSearchQuery(query);
  };

  if (loading) {
    return <Loading />;
  }

  return (
    <div className="animate__animated animate__fadeIn">
      <h1 className="text-3xl font-bold text-gray-800 mb-6">Todas las Actividades</h1>
      <SearchBar onSearch={handleSearch} />
      <ActivityList activities={activities} />
    </div>
  );
};

export default Activities;