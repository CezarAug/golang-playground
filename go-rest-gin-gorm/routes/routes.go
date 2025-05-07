package routes

import (
	"github.com/gin-gonic/gin"
	"study.co.jp/go-rest-gin-gorm/controllers"
)

func HandleRequests() {
	router := gin.Default()
	// Default port: 8080
	// Adding localhost to avoid windows asking everytime for firewall permissions

	router.GET("/students", controllers.GetAllStudents)
	router.GET("/students/:id", controllers.GetStudentByID)
	router.PATCH("/students/:id", controllers.UpdateStudent)
	router.DELETE("/students/:id", controllers.DeleteStudent)
	router.POST("/students", controllers.CreateStudent)

	router.GET("/students/number/:number", controllers.GetStudentByNumber)

	router.GET("/salute/:name", controllers.Salute)
	router.Run("localhost:8080")
}
