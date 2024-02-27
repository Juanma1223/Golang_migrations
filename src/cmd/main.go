package main

import (
	"flag"
	"fmt"
	"go-migrations/src/configHandler"
	dbhelper "go-migrations/src/dbHelper"
	filescreator "go-migrations/src/filesCreator"
	sqlparser "go-migrations/src/parser"
	sqlexec "go-migrations/src/sqlExec"
	versionhandler "go-migrations/src/versionHandler"
	"runtime"
)

func main() {
	fmt.Println("Golang Migrations CLI v0.8")

	// Get flag arguments
	dir := flag.String("dir", "./doc/db/migrations", "Directorory where migrations are located")

	// Database arguments
	dbUser := flag.String("u", "", "Data base username")
	dbPassword := flag.String("p", "", "Data base password")
	dbHost := flag.String("h", "", "Data base host")
	dbPort := flag.String("P", "", "Data base port")
	db := flag.String("d", "", "Data base name")
	changeName := flag.Bool("change", false, "Change database name")
	env := flag.String("env", "", "Select env number without the need to write it")

	// Migrations arguments
	revert := flag.Bool("revert", false, "If true, revert migrations")
	steps := flag.Int("steps", 0, "Number of steps for migrations")
	version := flag.Bool("version", false, "Get database version")
	forceVersion := flag.Int("force", 0, "New forced database version")
	fix := flag.Bool("fix", false, "Fix migration files prefix version")

	// Creation arguments
	newMigrationName := flag.String("create", "", "Create new migration with specific name")
	parseDir := flag.String("directory", "", "Parse new migrations based on content in directory")
	parseMigration := flag.String("parse", "", "Parse new migration from go struct on specified directory")
	flag.Parse()

	/*
		This arguments don't need user prompt, if the flag is set ignore any other flag or user input
	*/

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

	// Create migrations based on all models in directory recursively
	if *parseDir != "" {
		sqlparser.ParseDir(*parseDir, *dir)
		return
	}

	if *parseMigration != "" {
		sqlparser.ParseSql(*parseMigration, *dir)
		return
	}

	/*
		This arguments need user prompt
	*/

	// used to get the current path, if the terminal is in root path, the json file can be there
	_, path, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}

	if *env == "" {
		*env = configHandler.GetEnviromentFromUser()
	}
	settedFlags := configHandler.CheckFlags(db, dbUser, dbHost, dbPort, *env, path)
	// db is setted manually by the user and so is the password
	if *db == "" {
		*db = configHandler.GetDbFromUser()
	}
	if *dbPassword == "" {
		*dbPassword = configHandler.GetDbPassword()
	}
	// Set database parameters collected by CLI flags
	dbhelper.SetParams(*db, settedFlags.Username, *dbPassword, settedFlags.Host, settedFlags.Port)

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

	// Change database name
	if *changeName {
		configHandler.ChangeDbDefaultNameByEnviroment(*db)
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
