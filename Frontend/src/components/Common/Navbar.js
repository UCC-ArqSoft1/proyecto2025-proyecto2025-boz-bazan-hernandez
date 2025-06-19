// Navbar.js - Versión actualizada con clases CSS modernas
import React, { useContext, useState } from 'react';
import { Link, useHistory } from 'react-router-dom';
import { AuthContext } from '../../context/AuthContext';
import './Navbar.css'; // Importar los estilos

const Navbar = () => {
  const { user, logout, isAdmin } = useContext(AuthContext);
  const history = useHistory();
  const [mobileMenuOpen, setMobileMenuOpen] = useState(false);

  const handleLogout = () => {
    logout();
    history.push('/');
    setMobileMenuOpen(false);
  };

  const toggleMobileMenu = () => {
    setMobileMenuOpen(!mobileMenuOpen);
  };

  const closeMobileMenu = () => {
    setMobileMenuOpen(false);
  };

  return (
    <nav className="navbar">
      <div className="navbar-container">
        <Link to="/" className="navbar-brand" onClick={closeMobileMenu}>
          GymApp
        </Link>
        
        {/* Desktop Navigation */}
        <div className="navbar-nav">
          <Link to="/activities" className="nav-link">
            🏃‍♂️ Actividades
          </Link>
          
          {user ? (
            <div className="user-menu">
              <Link to="/my-activities" className="nav-link">
                📋 Mis Actividades
              </Link>
              
              {isAdmin() && (
                <Link to="/admin/activities" className="admin-link">
                  Admin
                </Link>
              )}
              
              <div className="user-info">
                <div className="user-avatar">
                  {user.name ? user.name.charAt(0).toUpperCase() : 'U'}
                </div>
                <span>¡Hola, {user.name || 'Usuario'}!</span>
              </div>
              
              <button onClick={handleLogout} className="logout-btn">
                🚪 Cerrar Sesión
              </button>
            </div>
          ) : (
            <Link to="/login" className="login-btn">
              🔐 Iniciar Sesión
            </Link>
          )}
        </div>

        {/* Mobile Menu Toggle */}
        <button 
          className="mobile-menu-toggle"
          onClick={toggleMobileMenu}
          aria-label="Toggle mobile menu"
        >
          <span></span>
          <span></span>
          <span></span>
        </button>
      </div>

      {/* Mobile Menu */}
      <div className={`mobile-menu ${mobileMenuOpen ? 'active' : ''}`}>
        <div className="mobile-nav-links">
          <Link 
            to="/activities" 
            className="nav-link"
            onClick={closeMobileMenu}
          >
            🏃‍♂️ Actividades
          </Link>
          
          {user && (
            <Link 
              to="/my-activities" 
              className="nav-link"
              onClick={closeMobileMenu}
            >
              📋 Mis Actividades
            </Link>
          )}
          
          {user && isAdmin() && (
            <Link 
              to="/admin/activities" 
              className="admin-link"
              onClick={closeMobileMenu}
            >
              👑 Admin
            </Link>
          )}
        </div>
        
        <div className="mobile-user-actions">
          {user ? (
            <>
              <div className="user-info">
                <div className="user-avatar">
                  {user.name ? user.name.charAt(0).toUpperCase() : 'U'}
                </div>
                <span>¡Hola, {user.name || 'Usuario'}!</span>
              </div>
              <button onClick={handleLogout} className="logout-btn">
                🚪 Cerrar Sesión
              </button>
            </>
          ) : (
            <Link 
              to="/login" 
              className="login-btn"
              onClick={closeMobileMenu}
            >
              🔐 Iniciar Sesión
            </Link>
          )}
        </div>
      </div>
    </nav>
  );
};

export default Navbar;