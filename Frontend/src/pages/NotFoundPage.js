import React from 'react';
import { Link } from 'react-router-dom';

const NotFoundPage = () => {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gray-50 text-center px-4">
      <h1 className="text-6xl font-bold text-indigo-600 mb-4">404</h1>
      <h2 className="text-2xl font-semibold text-gray-800 mb-2">Página no encontrada</h2>
      <p className="text-gray-600 mb-6">Lo sentimos, la página que estás buscando no existe.</p>
      <Link
        to="/"
        className="bg-indigo-600 text-white px-6 py-3 rounded-md hover:bg-indigo-700 transition duration-300"
      >
        Volver al inicio
      </Link>
    </div>
  );
};

export default NotFoundPage;