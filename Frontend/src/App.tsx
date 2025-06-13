import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './Home';
import ActivityDetailPage from './ActivityDetail';
import './App.css';

function App() {
    return (
        <Router>
            <div className="App">
                <header className="app-header">
                    <div className="app-header-content">
                        <h1 className="app-title">💪 GymApp</h1>
                        <p className="app-subtitle">Tu gimnasio digital</p>
                    </div>
                </header>

                <main className="app-main">
                    <Routes>
                        <Route path="/" element={<Home />} />
                        <Route path="/actividad/:id" element={<ActivityDetailPage />} />
                        <Route path="*" element={
                            <div className="not-found">
                                <h2>Página no encontrada</h2>
                                <p>La página que buscas no existe.</p>
                                <a href="/">Volver al inicio</a>
                            </div>
                        } />
                    </Routes>
                </main>

                <footer className="app-footer">
                    <p>&copy; 2024 GymApp - Sistema de Gestión de Actividades Deportivas</p>
                </footer>
            </div>
        </Router>
    );
}

export default App;