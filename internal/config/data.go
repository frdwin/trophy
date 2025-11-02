package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	Fuzzy    string `json:"fuzzy"`
	Terminal string `json:"terminal"`
}

func Parse() (Data, error) {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return Data{}, fmt.Errorf("error getting user's home directory: %w", err)
	}

	configContentBytes, err := os.ReadFile(fmt.Sprintf("%s/%s", userHomeDir, ".config/trophy/config.json"))
	if err != nil {
		return Data{}, fmt.Errorf("error opening config file: %w", err)
	}

	var dat Data
	if err = json.Unmarshal(configContentBytes, &dat); err != nil {
		return Data{}, fmt.Errorf("error unmarshaling config struct: %w", err)
	}

	return dat, nil
}
