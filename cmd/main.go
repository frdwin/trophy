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
	"sort"
	"strings"

	"github.com/frdwin/trophy/internal/apps"
	"github.com/frdwin/trophy/internal/config"
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
		"/usr/bin/ghostty -e",
		"The terminal command of your choice to open terminal apps.",
	)

	flag.Parse()

	if dat, err := config.Parse(); err == nil {
		fuzzyFlag = dat.Fuzzy
		termFlag = dat.Terminal
	}
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
	sort.Strings(appNames)

	var stdout bytes.Buffer
	parsedFuzzy := strings.Split(fuzzyFlag, " ")
	fuzzy := exec.Command(parsedFuzzy[0], parsedFuzzy[1:]...)
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

	app.Exec(termFlag)
}
