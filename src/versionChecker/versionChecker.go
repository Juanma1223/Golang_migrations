package versionchecker

import (
	dbhelper "go-migrations/src/dbHelper"
	"log"
	"strings"
)

// Get current migrations version on database
func GetCurrentVersion() int {
	db, err := dbhelper.GetDB()
	if err != nil {
		log.Fatal(err)
	}
	// defer db.Close()

	stmt, err := db.Prepare(`SELECT
								curr_version
							FROM
								migrations_version`)
	// Check if versions hasn't been initialized yet
	if strings.Contains(err.Error(), "doesn't exist") {
		CreateVersionTable()
		return 1
	}
	if err != nil {
		log.Fatal(err)
	}

	var version int
	err = stmt.QueryRow().Scan(&version)
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
	defer db.Close()
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS
						migrations_version(id INT PRIMARY KEY AUTO_INCREMENT, 
							curr_version INT DEFAULT 1
							)`)
	if err != nil {
		log.Fatal(err)
	}
}

func SetNewVersion() {

}
