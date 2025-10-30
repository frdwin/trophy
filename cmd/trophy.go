package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func main() {
	appFileNames, err := getAppFileNames()
	if err != nil {
		log.Fatalln(err)
	}

	appList, err := parseAppList(appFileNames)
	if err != nil {
		log.Fatalln(err)
	}

	var appNames []string
	for _, app := range appList {
		appNames = append(appNames, app.name)
	}

	var stdout bytes.Buffer
	fuzzyFinder := exec.Command("/usr/bin/sk")
	fuzzyFinder.Stdin = strings.NewReader(strings.Join(appNames, "\n"))
	fuzzyFinder.Stdout = &stdout

	err = fuzzyFinder.Run()
	if err != nil {
		log.Fatalf("Error starting fuzzy finder: %s\n", err)
	}

	parsedStdout := strings.TrimRight(stdout.String(), "\n")
	app, err := appList.getApp(parsedStdout)
	if err != nil {
		log.Fatalf("Error getting chosen app: %s\n", err)
	}

	execPath, err := exec.LookPath(strings.Trim(app.cmd, " "))
	if err != nil {
		log.Fatalf("Error finding chosen app's path: %s\n", err)
	}

	env := os.Environ()
	args := []string{app.cmd}
	attr := &os.ProcAttr{
		Env: env,
		Sys: &syscall.SysProcAttr{
			Setpgid: true,
		},
		Files: []*os.File{nil, nil, nil},
	}

	_, err = os.StartProcess(execPath, args, attr)
	if err != nil {
		log.Fatalf("Error initializing chosen app: %s\n", err)
	}
}
