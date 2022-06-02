package controller

import (
	"interview/models"
	"interview/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetStudentById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id must be number"})
		return
	}
	sqlConnInterface, _ := c.Get("sqlConnection")
	st, err := services.GetStudentByID(sqlConnInterface, id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, st)
	}
}
func GetStudent(c *gin.Context) {
	sqlConnInterface, _ := c.Get("sqlConnection")

	students, err := services.GetStudent(sqlConnInterface)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, students)
	}

}
func PostStudent(c *gin.Context) {
	sqlConnInterface, _ := c.Get("sqlConnection")
	var st models.Student
	if err := c.BindJSON(&st); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	// tientv commit this code
	id, err := services.PostStudent(sqlConnInterface, st)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		st.Id = int(id)
		c.IndentedJSON(http.StatusCreated, st)
	}
}
func SearchStudent(c *gin.Context) {
	name := c.Param("fullname")
	sqlConnInterface, _ := c.Get("sqlConnection")
	// tientv commit this code
	st, err := services.SearchStudentByName(sqlConnInterface, name)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		if st == nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User hasn't found"})
			return
		}
		c.IndentedJSON(http.StatusCreated, st)
	}
}
