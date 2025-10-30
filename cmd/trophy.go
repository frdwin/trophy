package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
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
	fmt.Println(parsedStdout)
}
