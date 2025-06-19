// Footer.js - Versión actualizada con clases CSS modernas
import React from 'react';
import './Footer.css'; // Importar los estilos

const Footer = () => {
  const currentYear = new Date().getFullYear();

  return (
    <footer className="footer">
      <div className="footer-container">
        <div className="footer-content">
          {/* Información de la empresa */}
          <div className="footer-section company-info">
            <div className="company-logo">GymApp</div>
            <p className="company-description">
              Tu plataforma favorita para descubrir y participar en actividades deportivas 
              y de bienestar. Conectamos personas apasionadas por el fitness y el deporte.
            </p>
            <div className="company-stats">
              <div className="stat-item">
                <span className="stat-number">1000+</span>
                <span className="stat-label">Usuarios</span>
              </div>
              <div className="stat-item">
                <span className="stat-number">50+</span>
                <span className="stat-label">Actividades</span>
              </div>
              <div className="stat-item">
                <span className="stat-number">4.8★</span>
                <span className="stat-label">Rating</span>
              </div>
            </div>
          </div>

          {/* Enlaces rápidos */}
          <div className="footer-section">
            <h3>Enlaces Rápidos</h3>
            <ul className="footer-links">
              <li><a href="/">🏠 Inicio</a></li>
              <li><a href="/activities">🏃‍♂️ Actividades</a></li>
              <li><a href="/about">ℹ️ Acerca de</a></li>
              <li><a href="/contact">📞 Contacto</a></li>
              <li><a href="/help">❓ Ayuda</a></li>
            </ul>
          </div>

          {/* Categorías */}
          <div className="footer-section">
            <h3>Categorías</h3>
            <ul className="footer-links">
              <li><a href="/activities?category=fitness">💪 Fitness</a></li>
              <li><a href="/activities?category=yoga">🧘‍♀️ Yoga</a></li>
              <li><a href="/activities?category=deportes">⚽ Deportes</a></li>
              <li><a href="/activities?category=natacion">🏊‍♂️ Natación</a></li>
              <li><a href="/activities?category=baile">💃 Baile</a></li>
            </ul>
          </div>

          {/* Newsletter y redes sociales */}
          <div className="footer-section">
            <div className="newsletter">
              <h3>📧 Newsletter</h3>
              <p>Recibe las últimas actividades y ofertas especiales</p>
              <div className="newsletter-form">
                <input 
                  type="email" 
                  className="newsletter-input" 
                  placeholder="tu@email.com"
                />
                <button className="newsletter-btn">Suscribirse</button>
              </div>
            </div>
            
            <div className="social-links">
              <a href="#" className="social-link facebook" aria-label="Facebook">
                📘
              </a>
              <a href="#" className="social-link instagram" aria-label="Instagram">
                📷
              </a>
              <a href="#" className="social-link twitter" aria-label="Twitter">
                🐦
              </a>
              <a href="#" className="social-link linkedin" aria-label="LinkedIn">
                💼
              </a>
            </div>
          </div>
        </div>

        <div className="footer-bottom">
          <p className="footer-copyright">
            © {currentYear} GymApp - Todos los derechos reservados
          </p>
          <p className="footer-tech">
            Desarrollado con
            <span className="tech-stack">
              ⚛️ React
            </span>
            <span className="tech-stack">
              🎨 CSS3
            </span>
            <span className="tech-stack">
              ✨ Amor
            </span>
          </p>
        </div>
      </div>
    </footer>
  );
};

export default Footer;