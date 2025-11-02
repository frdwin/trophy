/*
 * Copyright (c) 2025 Frederico Winter
 * https://github.com/frdwin/trophy
 * frederico@diaswinter.com.br
 */

package apps

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
)

// The App struct represents an application with the following fields:
//
// Name (string): The display name of the application.
// Cmd (string): The command to execute when the application is launched.
// Term (bool): A flag indicating whether the application should be opened in a terminal.
// noDisplay (bool): A flag indicating whether the application should be displayed in menus.
// fname (string): The filename of the .desktop file from which this app was parsed.
type App struct {
	Name      string
	Cmd       string
	Term      bool
	noDisplay bool
	fname     string
}

// AppList is defined as a slice of App structs, providing a convenient way to
// handle multiple applications.
type AppList []App

// GetFileNames retrieves the file paths of all .desktop application files from system
// and user directories.
//
// Returns: a slice of strings containing the file paths of the .desktop files, or
// an error if any issues occur accessing the directories.
func GetFileNames() ([]string, error) {
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

// parseApp parses a given .desktop file and returns an App struct.
//
// Parameters:
// - filename (string): The path to the .desktop file to be parsed.
//
// Returns: an App instance filled with data from the specified file, or
// an error if the file cannot be read.
func parseApp(filename string) (App, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return App{}, fmt.Errorf("error opening application file: %s", filename)
	}
	newApp := App{fname: filename}
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "Name=") && newApp.Name == "" {
			newApp.Name = line[5:]
		}
		if strings.HasPrefix(line, "Exec=") && newApp.Cmd == "" {
			cmd := strings.Split(line[5:], "%")
			newApp.Cmd = cmd[0]
		}
		if strings.HasPrefix(line, "Terminal=") && line[9:] == "true" {
			newApp.Term = true
		}
		if strings.HasPrefix(line, "NoDisplay=") && line[10:] == "true" {
			newApp.noDisplay = true
		}
	}
	return newApp, nil
}

// ParseFileNames processes multiple .desktop filenames and returns an AppList.
//
// Parameters:
// - appFileNames ([]string): A slice of paths to .desktop files.
//
// Returns: an AppList containing parsed applications that are intended for display,
// excluding those marked with NoDisplay=true. An error may be returned if parsing issues arise.

func ParseFileNames(appFileNames []string) (AppList, error) {
	var list AppList
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

// GetApp retrieves an application by name from the AppList.
//
// Parameters:
// - name (string): The name of the application to find.
//
// Returns: the corresponding App if found, or an error indicating that the
// application does not exist in the list.

func (appList *AppList) GetApp(name string) (App, error) {
	for _, app := range *appList {
		if app.Name == name {
			return app, nil
		}
	}
	return App{}, errors.New("app not found")
}

func (app *App) Exec(tf string) {
	baseCmd := strings.Split(app.Cmd, " ")[0]
	execPath, err := exec.LookPath(strings.Trim(baseCmd, " "))
	if err != nil {
		log.Fatalf("Error finding chosen app's path: %s\n", err)
	}

	env := os.Environ()
	args := strings.Split(app.Cmd, " ")
	attr := &os.ProcAttr{
		Env: env,
		Sys: &syscall.SysProcAttr{
			Setpgid: true,
		},
		Files: []*os.File{nil, nil, nil},
	}

	if app.Term {
		termExecCmd := append(strings.Split(tf, " "), execPath)
		log.Println(termExecCmd)
		_, err = os.StartProcess(termExecCmd[0], termExecCmd, attr)
		if err != nil {
			log.Fatalf("Error initializing terminal app: %s\n", err)
		}
	} else {
		_, err = os.StartProcess(execPath, args, attr)
		if err != nil {
			log.Fatalf("Error initializing chosen app: %s\n", err)
		}
	}
}
