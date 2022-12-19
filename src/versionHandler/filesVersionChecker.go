package versionhandler

import (
	"fmt"
	"strconv"
	"strings"
)

// Returns file name version
func GetFileVersion(fileName string) int {
	filePrefix := fileName[0:6]
	fileVersion, err := strconv.Atoi(filePrefix)
	if err != nil {
		fmt.Println("Error on file", fileName, ": File version not correctly set or file not created through CLI")
		fmt.Println("Use \"migration --create\" instead")
		return 0
	}
	return fileVersion
}

// Checks if a file has been created correctly or not
func CheckFileName(fileName, suffix string) bool {

	println(fileName, "_"+suffix)

	// Checks if it ends with up or down
	expectedSuffix := "_" + suffix + ".sql"
	if !strings.HasSuffix(fileName, expectedSuffix) {
		fmt.Println("Error on file", fileName, ": Expecting file name suffix:", expectedSuffix)
		return false
	}

	// Checks that the file has it's version at the beginning defined by the first 6 digits
	filePrefix := fileName[0:6]
	fileVersion, err := strconv.Atoi(filePrefix)
	if err != nil || fileVersion <= 0 {
		fmt.Println("Error on file", fileName, "incorrect version prefix, please fix file name")
	}

	return true
}
