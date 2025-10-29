package main

import (
	"fmt"
	"os"
)

func getAppFileNames() ([]os.DirEntry, error) {
	systemApps, err := os.ReadDir("/usr/share/applications")
	if err != nil {
		return nil, fmt.Errorf("error reading system's applications directory: %s", err)
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting user's home directory: %s", err)
	}

	userAppsPath := fmt.Sprintf("%s%s", userHomeDir, "/.local/share/applications")
	userApps, err := os.ReadDir(userAppsPath)
	if err != nil {
		return nil, fmt.Errorf("error reading user's applications directory: %s", err)
	}

	appFileNames := append(systemApps, userApps...)

	return appFileNames, nil
}
