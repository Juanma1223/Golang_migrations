package dbhelper

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Database string
var Username string
var Password string
var Host string
var Port string

func GetDB() (*sql.DB, error) {
	connection_url := Username + ":" + Password + "@tcp(" + Host + ":" + Port + ")"
	db, err := sql.Open("mysql", connection_url+"/"+Database)
	if err != nil {
		log.Fatal(err)
	}
	return db, err
}

// Parameters setter
func SetParams(database, username, password, host, port string) {
	Database = database
	Username = username
	Password = password
	Host = host
	Port = port
}
