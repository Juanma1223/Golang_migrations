package filescreator

import (
	versionchecker "go-migrations/src/versionChecker"
	"log"
	"os"
	"strconv"
)

func CreateNewMigration(name, directory string) {
	// Create directory hierarchy
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	// Get last migration number and concatenate it with name
	currentVersion := versionchecker.GetCurrentVersion()
	zerosQuantity := 6 - len(strconv.Itoa(currentVersion))
	newFileName := ""
	for i := 0; i < zerosQuantity; i++ {
		newFileName = newFileName + "0"
	}
	newFileName = newFileName + name
	err = os.WriteFile(directory+newFileName, []byte(""), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
