package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lms/controllers"
	"lms/middlewares"
	"lms/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AddAllowHeaders("Authorization")
	r.Use(cors.New(config))

	r.GET("/profile/:username", middlewares.JwtAuthMiddleware(), controllers.CurrentUser)
	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	categories := r.Group("/category")
	categories.Use(middlewares.JwtAuthMiddleware())
	{
		books := categories.Group("/books")
		books.GET("/:id", models.GetBookByID)
		books.GET("", models.GetAllBooks)
	}

	//protected := r.Group("/api/admin")
	//protected.Use(middlewares.JwtAuthMiddleware())
	//protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
