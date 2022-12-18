package sqlexec

import (
	"fmt"
	dbhelper "go-migrations/src/dbHelper"
	versionhandler "go-migrations/src/versionHandler"
	"io/ioutil"
	"log"
	"os"
)

const UP = "up"
const DOWN = "down"

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

func ApplyMigrations(folder string, steps int) {
	// Files array will always be ordered sequentially
	// we can rely on this to apply migrations correctly

	currVersion := versionhandler.GetCurrentVersion()

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}
	// Default value was sent, apply all migrations left
	if steps == 0 {
		steps = len(files) - currVersion
	}

	firstIndex := 0

	// Up migrations are always after down migrations
	// Array starts at 0, so we need to shift index to apply migrations correctly
	firstIndex = (currVersion * 2) - 1

	// Check if we have applied all migrations or we haven't applied none
	if firstIndex >= len(files) {
		fmt.Println("All migrations have been applied!")
		return
	}

	for i := firstIndex; i <= firstIndex+steps; i += 2 {
		if i >= len(files) {
			fmt.Println("All migrations have been applied!")
			return
		}
		currFile := files[i].Name()
		// Check file name correctness
		if !versionhandler.CheckFileName(currFile, UP) {
			return
		}
		// Check file version correctness
		fileVersion := versionhandler.GetFileVersion(currFile)
		if fileVersion != currVersion {
			fmt.Println("Expecting file version:", currVersion, ", have", fileVersion)
			return
		}
		fmt.Println("Applying migration: ", currFile)
		query, err := os.ReadFile(folder + "/" + currFile)
		if err != nil {
			log.Fatal(err)
		}
		ApplyMigration(string(query))
		fmt.Println("Success!")
		currVersion = currVersion + 1
		versionhandler.SetNewVersion(currVersion)
	}
}

func RevertMigrations(folder string, steps int) {
	// Files array will always be ordered sequentially
	// we can rely on this to apply migrations correctly

	currVersion := versionhandler.GetCurrentVersion()

	files, err := ioutil.ReadDir(folder)
	if err != nil {
		log.Fatal(err)
	}

	firstIndex := 0

	// Up migrations are always after down migrations
	// Array starts at 0, so we need to shift index to apply migrations correctly
	firstIndex = ((currVersion - 1) * 2) - 2

	// Index on which we have to stop reverting
	bottom := 0

	if steps == 0 {
		// Default value was sent, apply all migrations left
		bottom = -2
	} else {
		bottom = firstIndex - steps*2
	}

	if firstIndex < 0 {
		fmt.Println("No migrations left to revert")
		return
	}

	for i := firstIndex; i > bottom; i -= 2 {
		if i < 0 {
			fmt.Println("No migrations left to revert")
			return
		}
		currFile := files[i].Name()
		// Check file name correctness
		if !versionhandler.CheckFileName(currFile, DOWN) {
			return
		}
		// Check file version correctness
		fileVersion := versionhandler.GetFileVersion(currFile)
		if fileVersion != (currVersion - 1) {
			fmt.Println("Expecting file version:", (currVersion - 1), ", have", fileVersion)
			return
		}
		fmt.Println("Reverting migration: ", currFile)
		query, err := os.ReadFile(folder + "/" + currFile)
		if err != nil {
			log.Fatal(err)
		}
		ApplyMigration(string(query))
		fmt.Println("Success!")
		currVersion = currVersion - 1
		versionhandler.SetNewVersion(currVersion)
	}
}
