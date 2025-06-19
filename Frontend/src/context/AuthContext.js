import React, { createContext, useState, useEffect } from 'react';
import { login as authLogin, logout as authLogout, getCurrentUser } from '../api/authService';

export const AuthContext = createContext();

export const AuthProvider = ({ children }) => {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const user = getCurrentUser();
    if (user) {
      setUser(user);
    }
    setLoading(false);
  }, []);

  const login = async (email, password) => {
    try {
      const { token, role } = await authLogin(email, password);
      const user = { email, role };
      localStorage.setItem('token', token);
      localStorage.setItem('user', JSON.stringify(user));
      setUser(user);
      return { success: true };
    } catch (error) {
      return { success: false, error };
    }
  };

  const logout = () => {
    authLogout();
    setUser(null);
  };

  const isAdmin = () => {
    return user?.role === 'administrador';
  };

  return (
    <AuthContext.Provider value={{ user, loading, login, logout, isAdmin }}>
      {children}
    </AuthContext.Provider>
  );
};