package defaultConfig

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Config struct {
	Name     string `json:"name"`
	Database string `json:"database"`
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

var CurrentPath string

func GetDefaultConfigByName(input, path string) Config {
	jsons := GetAllEnvsFromJson(path)
	var env = jsons[input]
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

func BoolChecker() bool {
	var answer string
	var check bool
	fmt.Scanf("%s", &answer)
	answer = strings.ToLower(answer)
	switch answer {
	case "y", "yes":
		check = true
	default:
		check = false
		fmt.Println("skipping...")
	}

	return check
}

func CheckFlags(dbName, dbUser, dbHost, dbPort *string, input, path string) Config {
	config := GetDefaultConfigByName(input, path)
	if *dbUser == "" {
		*dbUser = config.Username
	}
	if *dbHost == "" {
		*dbHost = config.Host
	}
	if *dbPort == "" {
		*dbPort = config.Port
	}
	if *dbName == "" {
		*dbName = config.Database
	}
	if config.Database == "" {
		fmt.Println("Database default name not setted, do you want to set it now? (y/n)")
		check := BoolChecker()
		if check {
			newName := ChangeDbDefaultNameByEnviroment(input)
			*dbName = newName
		} else {
			*dbName = GetDbFromUser()
		}
		return config
	}
	return config
}

func ChangeDbDefaultNameByEnviroment(selectedEnviroment string) string {
	var enviromentName string
	fmt.Print("Enter database name to be used as default: ")
	fmt.Scanf("%s", &enviromentName)
	enviromentName = strings.ToLower(enviromentName)
	jsons := GetAllEnvsFromJson(CurrentPath)
	newName := SaveConfig(jsons, selectedEnviroment, enviromentName)
	return newName
}

func GetAllEnvsFromJson(path string) map[string]Config {
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
	CurrentPath = path

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading the json file", err)
	}

	var data map[string]Config
	err = json.Unmarshal([]byte(byteValue), &data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	return data
}

func SaveConfig(envs map[string]Config, env, name string) string {
	for selectedEnv, mapEnv := range envs {
		if env == mapEnv.Name {
			mapEnv.Database = name
			envs[selectedEnv] = mapEnv
		}
	}
	data, err := json.Marshal(envs)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	path := pwd + "/migrationConf.json"
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return ""
	}
	fmt.Println("Database name changed successfully!")
	return name
}
