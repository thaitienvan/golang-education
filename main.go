package main

import (
	"interview/connection"
	"interview/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.Use(connection.PointToDB)
	router.GET("/students", controller.GetStudent)
	router.GET("/student/:id", controller.GetStudentById)
	router.POST("/student", controller.PostStudent)
	router.Run("localhost:8111")
}
