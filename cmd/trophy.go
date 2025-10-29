package main

import "fmt"

func main() {
	appFileNames, err := getAppFileNames()
	if err != nil {
		panic(err)
	}

	for _, file := range appFileNames {
		fmt.Println(file.Name())
	}
}
