package config

import (
	"encoding/json"
	"os"
)

type Config struct {
    DbURL           string `json:"db_url"`
    CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, formatError("Error getting the config file path", err)
	}
	data, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, formatError("ReadFile failed: ", err)
	}
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, formatError("Unmarshal failed: ", err)
	}
	return config, nil
}

func (c *Config)SetUser(name string) error{
	c.CurrentUserName = name
	err:=write(*c)
	if err != nil {
		return formatError("Failed setting user", err)
	}
	return nil
}

func getConfigFilePath() (string, error) {
	const gatorconfig = "/.gatorconfig.json"
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", formatError("Failed getting file Path", err)
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
