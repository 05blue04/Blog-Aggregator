package config

import (
	"fmt"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url   string
	Username string
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return path + configFileName, nil
}

func Write(cfg Config) error {

	home, err := getConfigFilePath()

	fmt.Println(home)

	if err != nil {
		return err
	}

	data, err := os.ReadFile(home)

	if err != nil {
		return err
	}

	fmt.Println(data)

	return nil
}
