// pages/Auth/RegisterPage.js
import React, { useState } from 'react';
import { useHistory } from 'react-router-dom';
import RegisterForm from '../../components/Auth/RegisterForm';
import { registerUser } from '../../api/authService';
import Loading from '../../components/Common/Loading';
import Alert from '../../components/Common/Alert';

const RegisterPage = () => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const history = useHistory();

  const handleRegister = async (formData, errorData) => {
    if (errorData) {
      setError(errorData.error);
      return;
    }

    try {
      setLoading(true);
      setError(null);
      await registerUser({
        nombre: formData.nombre,
        email: formData.email,
        password: formData.password,
        tipoUsuario: false, // Por defecto se registra como socio
      });
      history.push('/login');
    } catch (error) {
      setError(error.response?.data?.error || 'Error al registrar el usuario');
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
    <div className="min-h-screen flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
      <div className="max-w-md w-full space-y-8">
        <RegisterForm onSubmit={handleRegister} error={error} onClearError={clearError} />
      </div>
    </div>
  );
};

export default RegisterPage;