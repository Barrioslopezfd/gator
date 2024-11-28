package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}


func Read() Config {
	filePath, err := getConfigFilePath()
	if err != nil {
		log.Fatal(err)
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err, "ReadFile failed: ", err)
	}
	var jsonFile Config
	err = json.Unmarshal(data, &jsonFile)
	if err != nil {
		log.Fatal(err, "Unmarshal failed: ", err)
	}
	fmt.Println(jsonFile)
	return jsonFile
}

func (c Config)SetUser(name string){
	c.CurrentUserName = name
	write(c)
}

func getConfigFilePath() (string, error) {
	const gatorconfig = "/.gatorconfig.json"
	dirname, err := os.UserHomeDir()
	if err != nil {
		newErr := fmt.Sprint("Failed getting file Path: ", err)
		return "", errors.New(newErr)
	}
	return (dirname+gatorconfig), nil
}

func write(c Config) error {
	filePath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	
	newData, err := json.MarshalIndent(c, "", "   ")
	if err != nil {
		return err
	}
	
	err = os.WriteFile(filePath, newData, 0644)
	if err != nil{
		return err
	}

	return nil
}
