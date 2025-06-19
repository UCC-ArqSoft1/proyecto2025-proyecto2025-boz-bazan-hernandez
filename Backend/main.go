package main

import (
	"log"
	"os"

	"gym-backend/controllers"
	"gym-backend/utils" // Para utils.ConnectDatabase()

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("No se encontró archivo .env, usando variables del sistema")
	}

	//CONECTAR A LA BASE DE DATOS PRIMERO
	utils.ConnectDatabase() // Esta función hace la auto-migración

	// Configurar Gin
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

	// Middleware CORS
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Inicializar controladores (ya existentes)
	authController := controllers.NewAuthController()
	userController := controllers.NewUserController()
	activityController := controllers.NewActivityController()

	// Rutas (ya existentes + nueva ruta de registro)
	api := router.Group("/api")
	{
		api.POST("/login", authController.Login)
		api.POST("/registro", authController.Register) // NUEVA RUTA AGREGADA
		api.GET("/actividades", activityController.GetAllActivities)
		api.GET("/actividades/:id", activityController.GetActivityByID)
	}

	protected := api.Group("/")
	protected.Use(utils.AuthMiddleware())
	{
		protected.GET("/mis-actividades", userController.GetMyActivities)
		protected.POST("/inscribirse/:id", userController.EnrollInActivity)

		admin := protected.Group("/")
		admin.Use(utils.AdminMiddleware())
		{
			admin.POST("/actividades", activityController.CreateActivity)
			admin.PUT("/actividades/:id", activityController.UpdateActivity)
			admin.DELETE("/actividades/:id", activityController.DeleteActivity)
		}
	}

	// Endpoint de salud + base de datos
	router.GET("/health", func(c *gin.Context) {
		// Verificar conexión a BD
		sqlDB, err := utils.DB.DB()
		if err != nil {
			c.JSON(500, gin.H{"status": "error", "database": "disconnected"})
			return
		}

		if err := sqlDB.Ping(); err != nil {
			c.JSON(500, gin.H{"status": "error", "database": "ping_failed"})
			return
		}

		c.JSON(200, gin.H{"status": "ok", "database": "connected"})
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciando en puerto %s", port)
	log.Printf("Base de datos: MySQL conectada y migrada")
	log.Printf("Admin por defecto: admin@gym.com / admin123")
	log.Printf("Registro habilitado en: POST /api/registro")

	log.Fatal(router.Run(":" + port))
}
