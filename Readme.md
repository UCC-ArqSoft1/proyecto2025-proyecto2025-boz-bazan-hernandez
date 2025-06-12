# Sistema de Gestión de Actividades Deportivas

Backend desarrollado en Go para la gestión de actividades deportivas de un gimnasio.

## Tecnologías
- Go 1.21+ - Lenguaje de Programacion
- MySQL - Base de datos relacional
- GORM - ORM para base de datos
- Gin - Framework web HTTP
- JWT - autenticación stateless

## Estructura del Proyecto
- backend/
- clients/           Clientes HTTP externos (para uso futuro)
- controllers/       Controladores REST API
- dao/              Acceso a datos
- domain/           Modelos de dominio y DTOs
- services/         Lógica de negocio
- utils/            Utilidades (JWT, hash, DB, middleware)
- main.go           Punto de entrada del servidor
- go.mod            Dependencias del proyecto
