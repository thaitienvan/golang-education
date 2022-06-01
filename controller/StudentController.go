package controller

import (
	"database/sql"
	"interview/models"
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
	var student models.Student
	sqlConnInterface, _ := c.Get("sqlConnection")
	sqlConn := sqlConnInterface.(*sql.DB)
	row := sqlConn.QueryRow("select * from student where student_id = ?", id)

	if rowScan := row.Scan(&student.Id, &student.FullName, &student.BirthDay, &student.PhoneNum, &student.Email); rowScan != nil {
		if rowScan == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Student doesn't exits"})
		}
	} else {
		c.IndentedJSON(http.StatusOK, student)
	}

}
func GetStudent(c *gin.Context) {
	sqlConnInterface, _ := c.Get("sqlConnection")
	sqlConn := sqlConnInterface.(*sql.DB)
	rows, err := sqlConn.Query("select * from student")
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	var students []models.Student
	for rows.Next() {
		var st models.Student
		err := rows.Scan(&st.Id, &st.FullName, &st.BirthDay, &st.PhoneNum, &st.Email)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		students = append(students, st)
	}
	c.IndentedJSON(http.StatusOK, students)
}
func PostStudent(c *gin.Context) {
	sqlConnInterface, _ := c.Get("sqlConnection")
	sqlConn := sqlConnInterface.(*sql.DB)

	var st models.Student
	if err := c.BindJSON(&st); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	inResult, err := sqlConn.Exec("INSERT INTO student(fullname,birthday,phone_num,email) VALUES (?, ?, ?, ?)", st.FullName, st.BirthDay, st.PhoneNum, st.Email)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	id, err := inResult.LastInsertId()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"id": id})

}
