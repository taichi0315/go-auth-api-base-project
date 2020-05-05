package infrastructure

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	dbHost := "localhost"
	dbUser := "user"
	dbPassword := "password"
	dbName := "auth_api"

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPassword, dbHost, dbName)

	var err error
	DB, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Println("error")
	}
}
