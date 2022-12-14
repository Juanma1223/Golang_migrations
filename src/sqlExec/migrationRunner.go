package sqlexec

import (
	"fmt"
	dbhelper "go-migrations/src/dbHelper"
	versionhandler "go-migrations/src/versionHandler"
	"io/ioutil"
	"log"
	"os"
)

func ApplyMigration(query string) {
	db, err := dbhelper.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	// Execute generic sql statement
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func ApplyMigrations(folder string) {
	// Check migration current version
	currVersion := versionhandler.GetCurrentVersion()

	// Read directory and apply every single migration available
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	for i, file := range files {
		// Only apply migrations after current database version
		if !file.IsDir() && (i >= currVersion) {
			// Read query from file
			query, err := os.ReadFile(folder + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Applying migration: ", file.Name())
			ApplyMigration(string(query))
			fmt.Println("Success!")
			// Update new migration version
			versionhandler.SetNewVersion(currVersion + 1)
			currVersion += 1
		}
	}
}
