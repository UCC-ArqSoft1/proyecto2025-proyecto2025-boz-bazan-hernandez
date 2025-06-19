import React from 'react';
import { useContext } from 'react';
import { AuthContext } from '../../context/AuthContext';
import Alert from '../Common/Alert';

const ActivityDetail = ({ activity, onEnroll, error, onClearError }) => {
  const { user } = useContext(AuthContext);

  return (
    <div className="bg-white rounded-lg shadow-lg overflow-hidden animate__animated animate__fadeIn">
      {error && <Alert message={error} type="error" onClose={onClearError} />}
      
      <div className="md:flex">
        <div className="md:w-1/2">
          <div className="h-64 md:h-full bg-gray-200 overflow-hidden">
            {activity.fotoURL ? (
              <img src={activity.fotoURL} alt={activity.titulo} className="w-full h-full object-cover" />
            ) : (
              <div className="w-full h-full flex items-center justify-center bg-gray-100">
                <span className="text-gray-400">No hay imagen</span>
              </div>
            )}
          </div>
        </div>
        
        <div className="md:w-1/2 p-6">
          <h1 className="text-3xl font-bold text-gray-800 mb-4">{activity.titulo}</h1>
          <p className="text-gray-600 mb-6">{activity.descripcion}</p>
          
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4 mb-6">
            <div>
              <h3 className="text-sm font-semibold text-gray-500 uppercase">Categoría</h3>
              <p className="text-lg">{activity.categoria}</p>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-gray-500 uppercase">Día</h3>
              <p className="text-lg">{activity.dia}</p>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-gray-500 uppercase">Horario</h3>
              <p className="text-lg">{activity.horario}</p>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-gray-500 uppercase">Duración</h3>
              <p className="text-lg">{activity.duracion}</p>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-gray-500 uppercase">Cupo</h3>
              <p className="text-lg">{activity.cupo} personas</p>
            </div>
            <div>
              <h3 className="text-sm font-semibold text-gray-500 uppercase">Instructor</h3>
              <p className="text-lg">{activity.instructor}</p>
            </div>
          </div>
          
          {user && user.role === 'socio' && (
            <button
              onClick={onEnroll}
              className="w-full bg-green-600 text-white py-3 rounded-md hover:bg-green-700 transition duration-300"
            >
              Inscribirme en esta actividad
            </button>
          )}
        </div>
      </div>
    </div>
  );
};

export default ActivityDetail;