import React, { useState } from 'react';
import Alert from '../Common/Alert';

const ActivityForm = ({ initialData, onSubmit, error, onClearError }) => {
  const [formData, setFormData] = useState(
    initialData || {
      titulo: '',
      descripcion: '',
      categoria: '',
      instructor: '',
      dia: 'Lunes',
      horario: '09:00',
      duracion: '60 min',
      cupo: 10,
      fotoURL: '',
    }
  );

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(formData);
  };

  return (
    <div className="bg-white rounded-lg shadow-md p-6 animate__animated animate__fadeIn">
      {error && <Alert message={error} type="error" onClose={onClearError} />}
      
      <form onSubmit={handleSubmit}>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          <div>
            <label className="block text-gray-700 mb-2" htmlFor="titulo">
              Título*
            </label>
            <input
              type="text"
              id="titulo"
              name="titulo"
              value={formData.titulo}
              onChange={handleChange}
              className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
              required
            />
          </div>
          
          <div>
            <label className="block text-gray-700 mb-2" htmlFor="categoria">
              Categoría*
            </label>
            <input
              type="text"
              id="categoria"
              name="categoria"
              value={formData.categoria}
              onChange={handleChange}
              className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
              required
            />
          </div>
          
          <div>
            <label className="block text-gray-700 mb-2" htmlFor="instructor">
              Instructor*
            </label>
            <input
              type="text"
              id="instructor"
              name="instructor"
              value={formData.instructor}
              onChange={handleChange}
              className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
              required
            />
          </div>
          
          <div>
            <label className="block text-gray-700 mb-2" htmlFor="dia">
              Día de la semana*
            </label>
            <select
              id="dia"
              name="dia"
              value={formData.dia}
              onChange={handleChange}
              className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
              required
            >
              <option value="Lunes">Lunes</option>
              <option value="Martes">Martes</option>
              <option value="Miércoles">Miércoles</option>
              <option value="Jueves">Jueves</option>
              <option value="Viernes">Viernes</option>
              <option value="Sábado">Sábado</option>
              <option value="Domingo">Domingo</option>
            </select>
          </div>
          
          <div>
            <label className="block text-gray-700 mb-2" htmlFor="horario">
              Horario (HH:MM)*
            </label>
            <input
              type="time"
              id="horario"
              name="horario"
              value={formData.horario}
              onChange={handleChange}
              className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
              required
            />
          </div>
          
          <div>
            <label className="block text-gray-700 mb-2" htmlFor="duracion">
              Duración (minutos)*
            </label>
            <select
              id="duracion"
              name="duracion"
              value={formData.duracion}
              onChange={handleChange}
              className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
              required
            >
              <option value="30 min">30 minutos</option>
              <option value="45 min">45 minutos</option>
              <option value="60 min">1 hora</option>
              <option value="90 min">1.5 horas</option>
              <option value="120 min">2 horas</option>
            </select>
          </div>
          
          <div>
            <label className="block text-gray-700 mb-2" htmlFor="cupo">
              Cupo máximo*
            </label>
            <input
              type="number"
              id="cupo"
              name="cupo"
              min="1"
              value={formData.cupo}
              onChange={handleChange}
              className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
              required
            />
          </div>
          
          <div>
            <label className="block text-gray-700 mb-2" htmlFor="fotoURL">
              URL de la imagen
            </label>
            <input
              type="url"
              id="fotoURL"
              name="fotoURL"
              value={formData.fotoURL}
              onChange={handleChange}
              className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
            />
          </div>
          
          <div className="md:col-span-2">
            <label className="block text-gray-700 mb-2" htmlFor="descripcion">
              Descripción
            </label>
            <textarea
              id="descripcion"
              name="descripcion"
              rows="3"
              value={formData.descripcion}
              onChange={handleChange}
              className="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-indigo-500"
            ></textarea>
          </div>
        </div>
        
        <div className="mt-6">
          <button
            type="submit"
            className="w-full bg-indigo-600 text-white py-3 rounded-md hover:bg-indigo-700 transition duration-300"
          >
            {initialData ? 'Actualizar Actividad' : 'Crear Actividad'}
          </button>
        </div>
      </form>
    </div>
  );
};

export default ActivityForm;