package versionhandler

import (
	"fmt"
	dbhelper "go-migrations/src/dbHelper"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Get current migrations version on database
func GetCurrentVersion() int {
	db, err := dbhelper.GetDB()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(`SELECT
								version
							FROM
								schema_migrations`)
	// Check if versions hasn't been initialized yet
	if err != nil {
		if strings.Contains(err.Error(), "doesn't exist") {
			// Create database versioning table
			CreateVersionTable()
			return 1
		} else {
			log.Fatal(err)
		}
	}
	if err != nil {
		log.Fatal(err)
	}

	var version int
	err = stmt.QueryRow().Scan(&version)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
	return version
}

func CreateVersionTable() {
	db, err := dbhelper.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS
		schema_migrations
	(version BIGINT(20), 
		dirty BOOLEAN DEFAULT 0
		)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO
			schema_migrations
		VALUES (1, 0)`)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func SetNewVersion(version int) {
	db, err := dbhelper.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := db.Prepare(`UPDATE 
								schema_migrations 
							SET
								version = ?`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(version)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func GetLastMigrationIndex(dir string) int {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	if len(files) == 0 {
		return 0
	}
	// Files number must always be even, so we can always divide by two
	index := len(files) / 2
	return index
}

// If git merge has left files with the same version, use this to fix it
func FixFilesVersions(dir string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	if (len(files) % 2) != 0 {
		fmt.Println("Files quantity is incorrect, check if migration files have been deleted or renamed")
		return
	}
	expectedEndVersion := len(files) / 2
	for i := 0; i < expectedEndVersion; i++ {
		// Apply 0s padding to file version prefix
		zerosQuantity := 6 - len(strconv.Itoa(i+1))
		fileNamePrefix := ""
		for j := 0; j < zerosQuantity; j++ {
			fileNamePrefix = fileNamePrefix + "0"
		}
		fileNamePrefix = fileNamePrefix + strconv.Itoa(i+1)
		fileIndex := i * 2
		// Rename down migration
		currFileName := files[fileIndex].Name()
		newFileName := dir + "/" + fileNamePrefix + currFileName[6:]
		err := os.Rename(dir+"/"+currFileName, newFileName)
		if err != nil {
			log.Fatal(err)
		}
		// Rename up migration
		currFileName = files[fileIndex+1].Name()
		newFileName = dir + "/" + fileNamePrefix + currFileName[6:]
		err = os.Rename(dir+"/"+currFileName, newFileName)
		if err != nil {
			log.Fatal(err)
		}
	}
}
