package main

import (
	"flag"
	"fmt"
	dbhelper "go-migrations/src/dbHelper"
	filescreator "go-migrations/src/filesCreator"
	sqlexec "go-migrations/src/sqlExec"
	versionhandler "go-migrations/src/versionHandler"
)

func main() {
	fmt.Println("Golang Migrations CLI v0.5")

	// Get flag arguments
	dir := flag.String("dir", "./doc/db/migrations", "Directorory where migrations are located")

	// Database arguments
	dbUser := flag.String("u", "root", "Data base username")
	dbPassword := flag.String("p", "root", "Data base password")
	dbHost := flag.String("h", "localhost", "Data base host")
	dbPort := flag.String("P", "3306", "Data base port")
	db := flag.String("d", "test", "Data base name")

	// Migrations arguments
	revert := flag.Bool("revert", false, "If true, revert migrations")
	steps := flag.Int("steps", 0, "Number of steps for migrations")
	version := flag.Bool("version", false, "Get database version")
	forceVersion := flag.Int("force", 0, "New forced database version")
	fix := flag.Bool("fix", false, "Fix migration files prefix version")

	// Creation arguments
	newMigrationName := flag.String("create", "", "Create new migration with specific name")

	flag.Parse()

	// Set database parameters collected by CLI flags
	dbhelper.SetParams(*db, *dbUser, *dbPassword, *dbHost, *dbPort)

	// Return version and ignore other flags
	if *version {
		fmt.Println("Current database version: ", versionhandler.GetCurrentVersion())
		return
	}

	// Set version and ignore all other flags
	if *forceVersion > 0 {
		versionhandler.SetNewVersion(*forceVersion)
		fmt.Println("Succesfully forced version ", *forceVersion)
		return
	}

	// Fix migration files prefix versions on a certain directory
	if *fix {
		fmt.Println("Fixing files versions")
		versionhandler.FixFilesVersions(*dir)
		return
	}

	// Create migrations
	if *newMigrationName != "" {
		filescreator.CreateNewMigration(*newMigrationName, *dir)
		return
	}

	// Apply or revert migrations
	if !*revert {
		sqlexec.ApplyMigrations(*dir, *steps)
		return
	} else {
		sqlexec.RevertMigrations(*dir, *steps)
		return
	}
}
