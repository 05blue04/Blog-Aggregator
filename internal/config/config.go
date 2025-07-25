package config

import (
	"encoding/json"
	"os"
)

const configFileName = "/.gatorconfig.json"

type Config struct {
	Db_url   string `json:"db_url"`
	Username string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	cfg.Username = userName
	return write(*cfg)
}

func getConfigFilePath() (string, error) {
	path, err := os.UserHomeDir()

	if err != nil {
		return "", err
	}

	return path + configFileName, nil
}

func Read() (Config, error) {

	home, err := getConfigFilePath()

	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(home)

	if err != nil {
		return Config{}, err
	}

	var cfg Config

	if err := json.Unmarshal(data, &cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}

func write(cfg Config) error {
	path_to_config, err := getConfigFilePath()

	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	os.WriteFile(path_to_config, data, 0644)

	return nil
}
