import React from 'react';
import ActivityCard from './ActivityCard';

const ActivityList = ({ activities }) => {
  if (activities.length === 0) {
    return (
      <div className="text-center py-10">
        <p className="text-gray-500 text-lg">No se encontraron actividades</p>
      </div>
    );
  }

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
      {activities.map((activity) => (
        <ActivityCard key={activity.id} activity={activity} />
      ))}
    </div>
  );
};

export default ActivityList;