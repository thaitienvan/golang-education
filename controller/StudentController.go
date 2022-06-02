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
	id, err := services.PostStudent(sqlConnInterface, st)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	} else {
		st.Id = int(id)
		c.IndentedJSON(http.StatusCreated, st)
	}
	// inResult, err := sqlConn.Exec("INSERT INTO student(fullname,birthday,phone_num,email) VALUES (?, ?, ?, ?)", st.FullName, st.BirthDay, st.PhoneNum, st.Email)
	// if err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }
	// id, err := inResult.LastInsertId()
	// if err != nil {
	// 	c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
	// 	return
	// }
	// c.IndentedJSON(http.StatusCreated, gin.H{"id": id})

}
func DeleteStudent(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"errormessage": err.Error()})
	}
	sqlConnInterface, _ := c.Get("sqlConnection")
	res, err := services.DeleteStudentById(sqlConnInterface, id)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if res {
		c.IndentedJSON(http.StatusOK, gin.H{"Message": "Delete succesfullt"})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"Message": "Student not found"})
	}
}
