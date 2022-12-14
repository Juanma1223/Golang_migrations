package main

import (
	"flag"
	"fmt"
	dbhelper "go-migrations/src/dbHelper"
	sqlexec "go-migrations/src/sqlExec"
)

func main() {
	fmt.Println("Golang Migrations CLI v0.1")

	// Get flag argument
	dir := flag.String("dir", "./doc/db/migrations", "Directorory where migrations are located")
	dbUser := flag.String("u", "root", "Data base username")
	dbPassword := flag.String("p", "root", "Data base password")
	dbHost := flag.String("h", "localhost", "Data base host")
	dbPort := flag.String("P", "3306", "Data base port")
	db := flag.String("d", "test", "Data base name")

	flag.Parse()

	// Set database parameters collected by CLI flags
	dbhelper.SetParams(*db, *dbUser, *dbPassword, *dbHost, *dbPort)

	sqlexec.ApplyAllMigrations(*dir)

}
