package versionhandler

import (
	dbhelper "go-migrations/src/dbHelper"
	"io/ioutil"
	"log"
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
	curr_version
	FROM
	migrations_version`)
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
	migrations_version
	(id INT PRIMARY KEY AUTO_INCREMENT, 
		curr_version INT DEFAULT 1
		)`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`INSERT INTO
		migrations_version
		VALUES (0, 0)`)
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
	migrations_version 
	SET
	curr_version = ?`)
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
	if len(files) == 0 {
		return 0
	}
	if err != nil {
		log.Fatal(err)
	}
	lastFile := files[len(files)-1]
	strIndex := lastFile.Name()[:6]
	index, err := strconv.Atoi(strIndex)
	if err != nil {
		log.Fatal(err)
	}
	return index
}
