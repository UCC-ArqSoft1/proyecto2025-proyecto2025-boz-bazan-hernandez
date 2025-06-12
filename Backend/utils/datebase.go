package utils

import (
	"fmt"
	"log"
	"os"

	"gym-backend/domain" // Importar tus modelos

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Variable global para la conexión
var DB *gorm.DB

func ConnectDatabase() {
	var err error

	// Construir DSN desde variables de entorno
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	// Configuración GORM
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // Ver queries SQL
	}

	// Conectar a MySQL
	DB, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}

	log.Println("✅ Conexión a la base de datos establecida")

	// 🔥 AUTO-MIGRACIÓN - GORM crea las tablas automáticamente
	err = DB.AutoMigrate(
		&domain.User{},        // Crea tabla 'users'
		&domain.Activity{},    // Crea tabla 'activities'
		&domain.Inscription{}, // Crea tabla 'inscripciones'
	)
	if err != nil {
		log.Fatal("Error en la migración: ", err)
	}

	log.Println("✅ Auto-migración completada - Tablas creadas/actualizadas")

	// Crear datos iniciales
	createInitialData()
}

// Crear usuario admin por defecto y datos de ejemplo
func createInitialData() {
	// Usuario administrador por defecto
	var adminCount int64
	DB.Model(&domain.User{}).Where("tipo_usuario = ?", true).Count(&adminCount)

	if adminCount == 0 {
		admin := domain.User{
			Nombre:       "Administrador",
			Email:        "admin@gym.com",
			PasswordHash: "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f", // "admin123" en SHA256
			TipoUsuario:  true,
			Activo:       true,
		}

		if err := DB.Create(&admin).Error; err != nil {
			log.Printf("Error creando admin: %v", err)
		} else {
			log.Println("✅ Usuario administrador creado - Email: admin@gym.com, Password: admin123")
		}
	}

	// Actividades de ejemplo (opcional)
	var activityCount int64
	DB.Model(&domain.Activity{}).Count(&activityCount)

	if activityCount == 0 {
		activities := []domain.Activity{
			{
				Titulo:         "Boxeo",
				Descripcion:    "Clase de boxeo para todos los niveles",
				Categoria:      "Cardio",
				Instructor:     "Pepe López",
				DiaSemana:      "Martes",
				Horario:        "18:00",
				Duracion:       60,
				CupoMaximo:     10,
				CupoDisponible: 10,
				Foto:           "https://example.com/boxeo.jpg",
				Activo:         true,
			},
			{
				Titulo:         "Funcional",
				Descripcion:    "Entrenamiento funcional de resistencia",
				Categoria:      "Funcional",
				Instructor:     "Juan Hernández",
				DiaSemana:      "Miércoles",
				Horario:        "12:00",
				Duracion:       45,
				CupoMaximo:     15,
				CupoDisponible: 15,
				Foto:           "https://example.com/funcional.jpg",
				Activo:         true,
			},
		}

		if err := DB.Create(&activities).Error; err != nil {
			log.Printf("Error creando actividades de ejemplo: %v", err)
		} else {
			log.Println("✅ Actividades de ejemplo creadas")
		}
	}
}

// Función para obtener conexión de BD (útil para otros paquetes)
func GetDB() *gorm.DB {
	return DB
}

// Cerrar conexión (para cleanup)
func CloseDatabase() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Printf("Error obteniendo SQL DB: %v", err)
		return
	}

	if err := sqlDB.Close(); err != nil {
		log.Printf("Error cerrando conexión: %v", err)
	} else {
		log.Println("✅ Conexión a base de datos cerrada")
	}
}
