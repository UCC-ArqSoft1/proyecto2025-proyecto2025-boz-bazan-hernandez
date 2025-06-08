package main

import (
	"log"
	"os"

	"gym-management-system/config"
	"gym-management-system/controllers"
	"gym-management-system/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar base de datos
	config.InitDB()

	// Configurar Gin
	r := gin.Default()

	// Configurar CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Rutas públicas
	public := r.Group("/api")
	{
		// Autenticación
		public.POST("/auth/login", controllers.Login)
		public.POST("/auth/register", controllers.Register)

		// Actividades públicas
		public.GET("/activities", controllers.GetActivities)
		public.GET("/activities/:id", controllers.GetActivityByID)
		public.GET("/activities/search", controllers.SearchActivities)
	}

	// Rutas protegidas para usuarios autenticados
	user := r.Group("/api/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.GET("/inscriptions", controllers.GetUserInscriptions)
		user.POST("/inscriptions", controllers.CreateInscription)
		user.GET("/profile", controllers.GetUserProfile)
	}

	// Rutas protegidas para administradores
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware())
	{
		admin.POST("/activities", controllers.CreateActivity)
		admin.PUT("/activities/:id", controllers.UpdateActivity)
		admin.DELETE("/activities/:id", controllers.DeleteActivity)
		admin.GET("/users", controllers.GetAllUsers)
	}

	// Obtener puerto del entorno o usar 8080 por defecto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor corriendo en puerto %s", port)
	log.Fatal(r.Run(":" + port))
}
