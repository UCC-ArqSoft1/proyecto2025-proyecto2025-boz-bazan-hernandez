package config

import (
	"fmt"
	"log"
	"os"

	"gym-management-system/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Obtener variables de entorno
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "")
	dbName := getEnv("DB_NAME", "gym_management")

	// Crear DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Conectar a la base de datos
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	log.Println("Conexión a MySQL establecida correctamente")

	// Auto-migrar los modelos
	err = DB.AutoMigrate(&models.User{}, &models.Activity{}, &models.Inscription{})
	if err != nil {
		log.Fatal("Error al migrar los modelos:", err)
	}

	log.Println("Migración de tablas completada")
}

// Función helper para obtener variables de entorno con valor por defecto
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetDB retorna la instancia de la base de datos
func GetDB() *gorm.DB {
	return DB
}
