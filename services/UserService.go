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
func GetStudent(iSQL interface{}) ([]models.Student, error) {
	sqlConn := iSQL.(*sql.DB)

	var students []models.Student
	stRows, err := sqlConn.Query("select * from student")
	if err != nil {
		return nil, errors.New(err.Error())
	}
	for stRows.Next() {
		var st models.Student
		stScan := stRows.Scan(&st.Id, &st.FullName, &st.BirthDay, &st.PhoneNum, &st.Email)
		if stScan != nil {
			return nil, errors.New(stScan.Error())
		}
		students = append(students, st)
	}
	return students, nil
}
