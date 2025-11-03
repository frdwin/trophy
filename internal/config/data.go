package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// The Data struct holds configuration settings.
type Data struct {
	Fuzzy    string `json:"fuzzy"`
	Terminal string `json:"terminal"`
}

// Parse reads the configuration file and returns a Data struct containing the settings.
//
// Returns:
// A Data instance populated with values from the configuration file.
// An error if any issues occur during file operations or JSON unmarshaling.
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
