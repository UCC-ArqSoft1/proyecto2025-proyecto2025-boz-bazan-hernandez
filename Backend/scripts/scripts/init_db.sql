-- Este archivo se ejecuta automáticamente cuando se crea el contenedor MySQL
CREATE DATABASE IF NOT EXISTS gym_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE gym_db;

-- GORM creará las tablas automáticamente, pero aquí hay ejemplos de datos adicionales
-- Datos de ejemplo (opcional - también se crean en Go)
INSERT IGNORE INTO users (id, nombre, email, password_hash, tipo_usuario, fecha_creacion, activo) VALUES 
(1, 'Administrador', 'admin@gym.com', 'ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f', TRUE, NOW(), TRUE),
(2, 'Usuario Demo', 'usuario@gym.com', 'ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f', FALSE, NOW(), TRUE);

-- Configuración MySQL para performance
SET GLOBAL innodb_buffer_pool_size = 134217728; -- 128MB
SET GLOBAL max_connections = 200;

SELECT 'Base de datos gym_db inicializada correctamente' as mensaje;
