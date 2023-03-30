package main

import (
	"flag"
	"fmt"
	dbhelper "go-migrations/src/dbHelper"
	"go-migrations/src/defaultConfig"
	filescreator "go-migrations/src/filesCreator"
	sqlparser "go-migrations/src/parser"
	sqlexec "go-migrations/src/sqlExec"
	versionhandler "go-migrations/src/versionHandler"
	"runtime"
)

func main() {
	fmt.Println("Golang Migrations CLI v0.5")

	// Get flag arguments
	dir := flag.String("dir", "./doc/db/migrations", "Directorory where migrations are located")

	// Database arguments
	dbUser := flag.String("u", "", "Data base username")
	dbPassword := flag.String("p", "", "Data base password")
	dbHost := flag.String("h", "", "Data base host")
	dbPort := flag.String("P", "", "Data base port")
	db := flag.String("d", "", "Data base name")
	changeName := flag.Bool("change", false, "Change database name")

	// Migrations arguments
	revert := flag.Bool("revert", false, "If true, revert migrations")
	steps := flag.Int("steps", 0, "Number of steps for migrations")
	version := flag.Bool("version", false, "Get database version")
	forceVersion := flag.Int("force", 0, "New forced database version")
	fix := flag.Bool("fix", false, "Fix migration files prefix version")

	// Creation arguments
	newMigrationName := flag.String("create", "", "Create new migration with specific name")
	parseMigration := flag.String("parse", "", "Parse new migration from go struct on specified directory")

	flag.Parse()
	// used to get the current path, if the terminal is in root path, the json file can be there
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	input := defaultConfig.GetEnviromentFromUser()
	settedFlags := defaultConfig.CheckFlags(db, dbUser, dbHost, dbPort, input, path)
	// db is setted manually by the user and so is the password
	if *db == "" {
		*db = defaultConfig.GetDbFromUser()
	}
	if *dbPassword == "" {
		*dbPassword = defaultConfig.GetDbPassword()
	}
	// Set database parameters collected by CLI flags
	dbhelper.SetParams(*db, settedFlags.Username, *dbPassword, settedFlags.Host, settedFlags.Port)

	// Change database name
	if *changeName {
		defaultConfig.ChangeDbDefaultNameByEnviroment(input)
	}

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

	if *parseMigration != "" {
		sqlparser.ParseSql(*parseMigration, *dir)
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
		filescreator.CreateNewMigration(*newMigrationName, *dir, "", "")
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
