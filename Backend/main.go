<<<<<<< HEAD
package main

import (
	"log"
	"os"

	"gym-backend/controllers"
	"gym-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Println("No se encontró archivo .env, usando variables del sistema")
	}

	// Conectar a la base de datos
	utils.ConnectDatabase()

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

	// Inicializar controladores
	authController := controllers.NewAuthController()
	userController := controllers.NewUserController()
	activityController := controllers.NewActivityController()

	// Rutas públicas
	api := router.Group("/api")
	{
		// Autenticación
		api.POST("/login", authController.Login)

		// Actividades (públicas para socios)
		api.GET("/actividades", activityController.GetAllActivities)
		api.GET("/actividades/:id", activityController.GetActivityByID)
	}

	// Rutas protegidas (requieren autenticación)
	protected := api.Group("/")
	protected.Use(utils.AuthMiddleware())
	{
		// Funcionalidades de socio
		protected.GET("/mis-actividades", userController.GetMyActivities)
		protected.POST("/inscribirse/:id", userController.EnrollInActivity)

		// Funcionalidades de administrador
		admin := protected.Group("/")
		admin.Use(utils.AdminMiddleware())
		{
			admin.POST("/actividades", activityController.CreateActivity)
			admin.PUT("/actividades/:id", activityController.UpdateActivity)
			admin.DELETE("/actividades/:id", activityController.DeleteActivity)
		}
	}

	// Endpoint de salud
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Obtener puerto del servidor
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciando en puerto %s", port)
	log.Fatal(router.Run(":" + port))
}
=======
package main

import (
	"log"
	"os"

	"gym-backend/controllers"
	"gym-backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se encontró archivo .env, usando variables del sistema")
	}

	utils.ConnectDatabase()

	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()

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

	authController := controllers.NewAuthController()
	userController := controllers.NewUserController()
	activityController := controllers.NewActivityController()

	api := router.Group("/api")
	{
		api.POST("/login", authController.Login)

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

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor iniciando en puerto %s", port)
	log.Fatal(router.Run(":" + port))
}
>>>>>>> e9f915f0d8d09355f0f2c17b2ed95dc1b1fad0ed
