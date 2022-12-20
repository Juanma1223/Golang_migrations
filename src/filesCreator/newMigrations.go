package filescreator

import (
	versionhandler "go-migrations/src/versionHandler"
	"log"
	"os"
	"strconv"
)

func CreateNewMigration(name, directory, upContent, downContent string) {
	// Create directory hierarchy
	err := os.MkdirAll(directory, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	// Get last migration number and concatenate it with name
	lastFileId := versionhandler.GetLastMigrationIndex(directory)
	zerosQuantity := 6 - len(strconv.Itoa(lastFileId+1))
	newFileName := ""
	// Apply 0s padding
	for i := 0; i < zerosQuantity; i++ {
		newFileName = newFileName + "0"
	}
	newFileName = newFileName + strconv.Itoa(lastFileId+1) + "_" + name
	err = os.WriteFile(directory+"/"+newFileName+"_up.sql", []byte(upContent), 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(directory+"/"+newFileName+"_down.sql", []byte(downContent), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
