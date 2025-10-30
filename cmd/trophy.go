package main

import (
	"fmt"
)

func main() {
	appFileNames, err := getAppFileNames()
	if err != nil {
		panic(err)
	}

	appList, err := parseAppList(appFileNames)
	if err != nil {
		panic(err)
	}

	for _, app := range appList {
		fmt.Printf("Name: %s\n", app.name)
		fmt.Printf("Command: %s\n", app.cmd)
		fmt.Printf("File Name: %s\n", app.fname)
		fmt.Println()
	}

	fmt.Println(len(appList))
}
