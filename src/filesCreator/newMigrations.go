package filescreator

import (
	versionhandler "go-migrations/src/versionHandler"
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
	lastFileId := versionhandler.GetLastMigrationIndex(directory)
	zerosQuantity := 6 - len(strconv.Itoa(lastFileId+1))
	newFileName := ""
	for i := 0; i < zerosQuantity; i++ {
		newFileName = newFileName + "0"
	}
	newFileName = newFileName + strconv.Itoa(lastFileId+1) + "_" + name
	err = os.WriteFile(directory+"/"+newFileName+".sql", []byte(""), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
