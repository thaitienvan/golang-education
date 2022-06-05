package main

import (
	"interview/connection"
	"interview/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(connection.PointToDB)
	// router.GET("/students", controller.GetStudent)
	// router.GET("/student/:id", controller.GetStudentById)
	// router.POST("/student", controller.PostStudent)
	// router.GET("/searchstudent/:fullname", controller.SearchStudent)
	// router.DELETE("/student/:id", controller.DeleteStudent)
	st := router.Group("/student")
	{
		st.GET("/", controller.GetStudent)
		st.GET("/:id", controller.GetStudentById)
		st.POST("/", controller.PostStudent)
		st.GET("/search/:fullname", controller.SearchStudent)
		st.DELETE("/:id", controller.DeleteStudent)
		st.GET("/searchbyyear/:year", controller.FindStudentByYear)
	}
	router.Run("localhost:8111")
}
