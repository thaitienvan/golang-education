package connection

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type errorInterface interface {
	Error() string
}

func handleError(errinterface errorInterface) {
	if errinterface != nil {
		log.Println(errinterface.Error())
	}
}
func PointToDB(c *gin.Context) {
	log.Println("Setup connection to DB")
	var sqlConn *sql.DB
	var err error
	sqlConn, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/education")
	handleError(err)
	log.Println("Checking the connectivity")
	pingerr := sqlConn.Ping()
	handleError(pingerr)
	c.Set("sqlConnection", sqlConn)

}
