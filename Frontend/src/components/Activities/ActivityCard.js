import React from 'react';
import { Link } from 'react-router-dom';

const ActivityCard = ({ activity }) => {
  return (
    <div className="bg-white rounded-lg shadow-md overflow-hidden transition-transform duration-300 hover:scale-105 animate__animated animate__fadeInUp">
      <div className="h-48 bg-gray-200 overflow-hidden">
        {activity.fotoURL ? (
          <img src={activity.fotoURL} alt={activity.titulo} className="w-full h-full object-cover" />
        ) : (
          <div className="w-full h-full flex items-center justify-center bg-gray-100">
            <span className="text-gray-400">No hay imagen</span>
          </div>
        )}
      </div>
      <div className="p-4">
        <h3 className="text-xl font-semibold text-gray-800 mb-2">{activity.titulo}</h3>
        <p className="text-gray-600 mb-1">
          <span className="font-medium">Horario:</span> {activity.horario}
        </p>
        <p className="text-gray-600 mb-3">
          <span className="font-medium">Profesor:</span> {activity.profesor}
        </p>
        <Link
          to={`/activities/${activity.id}`}
          className="inline-block bg-indigo-600 text-white px-4 py-2 rounded-md hover:bg-indigo-700 transition duration-300"
        >
          Ver detalles
        </Link>
      </div>
    </div>
  );
};

export default ActivityCard;