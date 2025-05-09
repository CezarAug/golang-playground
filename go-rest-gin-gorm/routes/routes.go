package routes

import (
	"github.com/gin-gonic/gin"
	"study.co.jp/go-rest-gin-gorm/controllers"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "study.co.jp/go-rest-gin-gorm/docs"
)

func HandleRequests() {
	router := gin.Default()
	// Default port: 8080

	// Template and static assets sources
	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	// Routes
	router.GET("/students", controllers.GetAllStudents)
	router.GET("/students/:id", controllers.GetStudentByID)
	router.PATCH("/students/:id", controllers.UpdateStudent)
	router.DELETE("/students/:id", controllers.DeleteStudent)
	router.POST("/students", controllers.CreateStudent)

	router.GET("/students/number/:number", controllers.GetStudentByNumber)

	router.GET("/salute/:name", controllers.Salute)

	//Templates
	router.GET("/index", controllers.ShowIndexPage)
	router.NoRoute(controllers.RouteNotFoud)

	// Swagger
	//TODO: Prepare a better versioning
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// Adding localhost to avoid windows asking everytime for firewall permissions
	router.Run("localhost:8080")
}
