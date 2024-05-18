package app

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

/* create MySQL connection or die */
func (app *App) createMySQLConnection() {
	mysqlHost := os.Getenv("GYM_MENTOR_DB_HOST")
	mysqlPort := os.Getenv("GYM_MENTOR_DB_PORT")
	mysqlUser := os.Getenv("GYM_MENTOR_DB_USER")
	mysqlPassword := os.Getenv("GYM_MENTOR_DB_PASSWORD")
	mysqlDatabaseName := os.Getenv("GYM_MENTOR_DB_NAME")

	connectionURI := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDatabaseName)
	connection, connectionError := sql.Open("mysql", connectionURI)
	if connectionError != nil {
		log.Fatalln("FATAL ERROR: MySQLConnection failed", connectionError)
	}
	log.Println("MySQL connection created")
	app.MySQL = connection

}
