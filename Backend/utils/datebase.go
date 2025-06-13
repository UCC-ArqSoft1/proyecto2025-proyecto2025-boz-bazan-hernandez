package utils

import (
	"fmt"
	"log"
	"os"
	"time"

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

	//  GORM crea las tablas automáticamente
	err = DB.AutoMigrate(
		&domain.User{},        // Crea tabla 'users'
		&domain.Activity{},    // Crea tabla 'activities'
		&domain.Inscription{}, // Crea tabla 'inscripciones'
	)
	if err != nil {
		log.Fatal("Error en la migración: ", err)
	}

	log.Println(" Auto-migración completada - Tablas creadas/actualizadas")

	// Crear datos iniciales
	createInitialData()
}

// Crear usuario admin por defecto y datos de ejemplo
func createInitialData() {
	// Usuario administrador por defecto - SIEMPRE recrear para asegurar hash correcto
	var admin domain.User
	result := DB.Where("email = ?", "admin@gym.com").First(&admin)

	if result.Error != nil || admin.ID == 0 {
		// Crear nuevo admin
		hashedPassword := HashPassword("admin123")

		admin = domain.User{
			Nombre:       "Administrador",
			Email:        "admin@gym.com",
			PasswordHash: hashedPassword,
			TipoUsuario:  true,
			Activo:       true,
		}

		if err := DB.Create(&admin).Error; err != nil {
			log.Printf("Error creando admin: %v", err)
		} else {
			log.Printf("✅ Usuario administrador creado - Email: admin@gym.com, Password: admin123")
			log.Printf("🔑 Hash generado: %s", hashedPassword)
		}
	} else {
		// Actualizar hash del admin existente
		hashedPassword := HashPassword("admin123")
		admin.PasswordHash = hashedPassword

		if err := DB.Save(&admin).Error; err != nil {
			log.Printf("Error actualizando admin: %v", err)
		} else {
			log.Printf("🔄 Admin actualizado con nuevo hash: %s", hashedPassword)
		}
	}

	// Actividades de ejemplo (opcional)
	var activityCount int64
	DB.Model(&domain.Activity{}).Count(&activityCount)

	if activityCount == 0 {
		horarioBoxeo, _ := time.Parse("2006-01-02 15:04:05", "2025-06-13 18:00:00")
		horarioFuncional, _ := time.Parse("2006-01-02 15:04:05", "2025-06-13 12:00:00")
		activities := []domain.Activity{
			{
				Titulo:         "Boxeo",
				Descripcion:    "Clase de boxeo para todos los niveles",
				Categoria:      "Cardio",
				Instructor:     "Pepe López",
				DiaSemana:      "Martes",
				Horario:        horarioBoxeo,
				Duracion:       60,
				CupoMaximo:     10,
				CupoDisponible: 10,
				Activo:         true,
			},
			{
				Titulo:         "Funcional",
				Descripcion:    "Entrenamiento funcional de resistencia",
				Categoria:      "Funcional",
				Instructor:     "Juan Hernández",
				DiaSemana:      "Miércoles",
				Horario:        horarioFuncional,
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
			log.Println(" Actividades de ejemplo creadas")
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
