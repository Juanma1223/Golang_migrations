package main

import (
	"flag"
	"fmt"
	dbhelper "go-migrations/src/dbHelper"
	filescreator "go-migrations/src/filesCreator"
	sqlexec "go-migrations/src/sqlExec"
)

func main() {
	fmt.Println("Golang Migrations CLI v0.1")

	// Get flag arguments
	dir := flag.String("dir", "./doc/db/migrations", "Directorory where migrations are located")
	// dir := "../../doc/db/migrations"
	dbUser := flag.String("u", "root", "Data base username")
	dbPassword := flag.String("p", "root", "Data base password")
	dbHost := flag.String("h", "localhost", "Data base host")
	dbPort := flag.String("P", "3306", "Data base port")
	db := flag.String("d", "test", "Data base name")
	newMigrationName := flag.String("create", "", "Create new migration with specific name")
	// newMigrationName := "testing"

	flag.Parse()

	// Set database parameters collected by CLI flags
	dbhelper.SetParams(*db, *dbUser, *dbPassword, *dbHost, *dbPort)

	if *newMigrationName == "" {
		sqlexec.ApplyMigrations(*dir)
	} else {
		filescreator.CreateNewMigration(*newMigrationName, *dir)
	}

}
