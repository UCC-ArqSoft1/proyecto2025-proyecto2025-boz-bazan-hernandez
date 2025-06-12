package utils

import (
	"fmt"
	"log"
	"os"

	"gym-backend/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	DB, err = gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos: ", err)
	}

	log.Println("Conexión a la base de datos establecida")

	err = DB.AutoMigrate(&domain.User{}, &domain.Activity{}, &domain.Inscription{})
	if err != nil {
		log.Fatal("Error en la migración: ", err)
	}

	log.Println("Migración completada")

	createDefaultAdmin()
}

func createDefaultAdmin() {
	var count int64
	DB.Model(&domain.User{}).Where("tipo_usuario = ?", true).Count(&count)

	if count == 0 {
		admin := domain.User{
			Nombre:       "Administrador",
			Email:        "admin@gym.com",
			PasswordHash: "ef92b778bafe771e89245b89ecbc08a44a4e166c06659911881f383d4473e94f",
			TipoUsuario:  true,
			Activo:       true,
		}

		result := DB.Create(&admin)
		if result.Error != nil {
			log.Printf("Error creando admin por defecto: %v", result.Error)
		} else {
			log.Println("Usuario administrador por defecto creado - Email: admin@gym.com, Password: admin123")
		}
	}
}
