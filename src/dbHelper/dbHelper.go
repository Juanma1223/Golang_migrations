package dbhelper

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	// When using local development uncomment this line of code with your own credentials
	connection_url := os.Getenv("DB")

	if connection_url == "" {
		connection_url = "root:root@tcp(localhost:3306)"
	}
	db, err := sql.Open("mysql", connection_url+"/test")
	db.SetMaxOpenConns(40)
	if err != nil {
		return nil, err
	}

	return db, err
}
