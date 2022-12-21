package sqlparser

import (
	"fmt"
	filescreator "go-migrations/src/filesCreator"
	"io/ioutil"
	"log"
	"strings"

	"github.com/iancoleman/strcase"
)

// Define Finite Automaton state
var currState int = 0

func ParseSql(fileDir, outputDir string) {
	rawFileContent, err := ioutil.ReadFile(fileDir)
	fileContent := string(rawFileContent)
	if err != nil {
		log.Fatal(err)
	}

	// Divide content by lines
	lines := strings.Split(fileContent, "\n")

	// Skip package line
	lines = lines[1:]

	// Clean comments and line jumps
	lines = cleanLines(lines)

	// Parse struct initial syntax skipping package line

	migrationName := strcase.ToSnake(parseStructInit(lines[0]))

	// Pop currentLine from slice
	lines = lines[1:]
	var tableFields string
	for _, line := range lines {
		if line == "}" {
			break
		}
		tableFields = tableFields + "\t" + parseColumn(line) + ",\n"
	}
	// Remove last comma and line jump
	tableFields = tableFields[:len(tableFields)-2]
	sqlQuery := "CREATE TABLE " + migrationName + "(\n" + tableFields + "\n)"

	migrationFileName := "create_" + migrationName + "_table"
	filescreator.CreateNewMigration(migrationFileName, outputDir, sqlQuery, "DROP TABLE "+migrationName)
}

// Removes white spaces and line jumps
func cleanLines(lines []string) []string {
	cleanLines := []string{}
	for _, line := range lines {
		isComment := strings.HasPrefix(line, "//")
		if line != "" && !isComment {
			cleanLines = append(cleanLines, line)
		}
	}
	return cleanLines
}

// Remove white spaces and comments
func lineCleaner(line []string) []string {
	cleanLine := []string{}
	for _, word := range line {
		// Ignore white spaces and line jumps
		if word != "" {
			cleanLine = append(cleanLine, word)
		}
	}
	return cleanLine
}

// Returns new migration name and validates if struct declaring syntax is valid
func parseStructInit(stringLine string) string {
	line := strings.Split(stringLine, " ")
	line = lineCleaner(line)

	if len(line) < 4 {
		fmt.Println("Error: Bad syntax on struct initialization")
		return ""
	}

	if line[0] == "type" {
		currState += 1
	} else {
		fmt.Println("Error: Expecting \"type\", found:", line[0])
		return ""
	}

	migrationName := line[1]

	if line[2] == "struct" {
		currState += 1
	} else {
		fmt.Println("Error: Expecting \"struct\", found:", line[2])
		return ""
	}

	if line[3] == "{" {
		currState += 1
	} else {
		fmt.Println("Error: Expecting \"}\", found:", line[3])
		return ""
	}
	return migrationName
}

func parseColumn(stringLine string) string {
	line := strings.Split(stringLine, " ")
	// Clean white spaces and comments
	line = lineCleaner(line)

	if len(line) < 3 {
		log.Fatal("Error: Column", stringLine, "syntax is incorrect")
		return ""
	}
	columnName := line[0]
	goDataType := line[1]

	var sqlDataType string

	// Parse sql data type
	switch goDataType {
	case "int":
		sqlDataType = "INT DEFAULT 0"
	case "string":
		sqlDataType = "VARCHAR(50) DEFAULT ''"
	case "bool":
		sqlDataType = "BOOLEAN DEFAULT 0"
	case "float":
		sqlDataType = "FLOAT DEFAULT 0"
	default:
		sqlDataType = "VARCHAR(50) DEFAULT ''"
	}

	// Check if column is primary key
	if columnName != "id" {
		return strcase.ToSnake(columnName) + " " + sqlDataType
	} else {
		return "id INT PRIMARY KEY AUTO_INCREMENT"
	}
}
