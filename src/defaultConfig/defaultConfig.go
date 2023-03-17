package defaultConfig

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	Database string `json:"database"`
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func GetDefaultConfig(input, path string) Config {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	path = pwd + "/migrationConf.json"
	jsonFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading the json file", err)
	}

	var data map[string]Config
	err = json.Unmarshal([]byte(byteValue), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return Config{}
	}
	var env = data[input]
	return env
}

func GetDbPassword() string {
	var password string
	fmt.Print("Enter password: ")
	fmt.Scanf("%s", &password)
	return password
}

func GetEnviromentFromUser() string {
	var env string
	fmt.Print("Enter enviroment: ")
	fmt.Scanf("%s", &env)
	env = strings.ToUpper(env)
	return env
}

func GetDbFromUser() string {
	var db string
	fmt.Print("Enter database: ")
	fmt.Scanf("%s", &db)
	return db
}

func CheckFlags(dbUser, dbHost, dbPort *string, input, path string) Config {
	config := GetDefaultConfig(input, path)
	if *dbUser == "" {
		*dbUser = config.Username
	}
	if *dbHost == "" {
		*dbHost = config.Host
	}
	if *dbPort == "" {
		*dbPort = config.Port
	}
	return config
}
