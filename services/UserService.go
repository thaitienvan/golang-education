package services

import (
	"database/sql"
	"errors"
	"fmt"
	"interview/models"
)

func GetStudentByID(iSQL interface{}, stID int) (models.Student, error) {
	sqlConn := iSQL.(*sql.DB)
	var st models.Student
	row := sqlConn.QueryRow("SELECT * FROM student where student_id=?", stID)
	if stScan := row.Scan(&st.Id, &st.FullName, &st.BirthDay, &st.PhoneNum, &st.Email); stScan != nil {
		if stScan == sql.ErrNoRows {
			return models.Student{}, errors.New("student doesn't exist")
		} else {
			return models.Student{}, fmt.Errorf("get student error %v", stScan)
		}
	} else {
		return st, nil
	}
}
