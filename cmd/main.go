// TODO:
// - open terminal apps in the terminal chosen by the user with -t flag
// - fix bug when opening commands with more than one arg (v.g., steam games)
// - add possibility to configure trophy with a config file
// - review README

/*
 * Copyright (c) 2025 Frederico Winter
 * https://github.com/frdwin/trophy
 * frederico@diaswinter.com.br
 */

package main

import (
	"bytes"
	"flag"
	"log"
	"os/exec"
	"strings"

	"github.com/frdwin/trophy/cmd/apps"
)

var fuzzyFlag, termFlag string

func init() {
	flag.StringVar(
		&fuzzyFlag,
		"f",
		"/usr/bin/sk",
		"The fuzzy finder application of your choice.",
	)

	flag.StringVar(
		&termFlag,
		"t",
		"/usr/bin/ghostty",
		"The terminal application of your choice.",
	)

	flag.Parse()
}

func main() {
	appFileNames, err := apps.GetFileNames()
	if err != nil {
		log.Fatalln(err)
	}

	appList, err := apps.ParseFileNames(appFileNames)
	if err != nil {
		log.Fatalln(err)
	}

	var appNames []string
	for _, app := range appList {
		appNames = append(appNames, app.Name)
	}

	var stdout bytes.Buffer
	fuzzy := exec.Command(fuzzyFlag)
	fuzzy.Stdin = strings.NewReader(strings.Join(appNames, "\n"))
	fuzzy.Stdout = &stdout

	err = fuzzy.Run()
	if err != nil {
		log.Fatalf("Error starting fuzzy finder: %s\n", err)
	}

	parsedStdout := strings.TrimRight(stdout.String(), "\n")
	app, err := appList.GetApp(parsedStdout)
	if err != nil {
		log.Fatalf("Error getting chosen app: %s\n", err)
	}

	app.Exec()
}
