package sqlexec

import (
	dbhelper "go-migrations/src/dbHelper"
	"io/ioutil"
	"log"
	"os"
)

func ApplyMigration(query string) {
	db, err := dbhelper.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Execute generic sql statement
	_, err = db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func ApplyAllMigrations(folder string) {
	// Check if version is correct

	// Read directory and apply every single migration available
	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if !file.IsDir() {
			// Read query from file
			query, err := os.ReadFile(folder + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			ApplyMigration(string(query))
		}
	}
}
