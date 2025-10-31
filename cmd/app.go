package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type app struct {
	name      string
	cmd       string
	noDisplay bool
	fname     string
}

type appList []app

func getAppFileNames() ([]string, error) {
	var sysAppsFileNames, userAppsFileNames []string

	sysAppsPathSlice := []string{"/usr/share/applications", "/usr/local/share/applications"}
	for _, sysAppsPath := range sysAppsPathSlice {
		_, err := os.Stat(sysAppsPath)
		if err == nil {
			sysApps, err := os.ReadDir(sysAppsPath)
			if err != nil {
				return nil, fmt.Errorf("error reading system's applications directory: %s", err)
			}

			for _, sysApp := range sysApps {
				appFileName := filepath.Join(sysAppsPath, sysApp.Name())
				if strings.HasSuffix(appFileName, ".desktop") {
					sysAppsFileNames = append(sysAppsFileNames, appFileName)
				}
			}
		}
	}

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("error getting user's home directory: %s", err)
	}

	userAppsPath := filepath.Join(userHomeDir, "/.local/share/applications")
	userApps, err := os.ReadDir(userAppsPath)
	if err != nil {
		return nil, fmt.Errorf("error reading user's applications directory: %s", err)
	}

	for _, userApp := range userApps {
		appFileName := filepath.Join(userAppsPath, userApp.Name())
		if strings.HasSuffix(appFileName, ".desktop") {
			userAppsFileNames = append(userAppsFileNames, appFileName)
		}
	}

	appFileNames := append(sysAppsFileNames, userAppsFileNames...)

	return appFileNames, nil
}

func parseApp(filename string) (app, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return app{}, fmt.Errorf("error opening application file: %s", filename)
	}
	newApp := app{fname: filename}
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Name=") && newApp.name == "" {
			newApp.name = line[5:]
		}
		if strings.HasPrefix(line, "Exec=") && newApp.cmd == "" {
			cmd := strings.Split(line[5:], "%")
			newApp.cmd = cmd[0]
		}
		if strings.HasPrefix(line, "NoDisplay=") && line[10:] == "true" {
			newApp.noDisplay = true
		}
	}
	return newApp, nil
}

func parseAppList(appFileNames []string) (appList, error) {
	var list appList
	for _, fname := range appFileNames {
		newApp, err := parseApp(fname)
		if err != nil {
			fmt.Printf("Error: could not parse %s\n", fname)
		}
		if !newApp.noDisplay {
			list = append(list, newApp)
		}
	}
	return list, nil
}

func (al *appList) getApp(name string) (app, error) {
	for _, app := range *al {
		if app.name == name {
			return app, nil
		}
	}
	return app{}, errors.New("app not found")
}
