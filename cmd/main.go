// - diferenciar os apps que precisam de terminal
// e abri-los em um escolhido pelo usuário
// - permitir configurar o trophy
// - adicionar a linha de comando: escolha do terminal e do fuzzy finder
// - consertar o bug nos aplicativos com argumentos (steam games)
package main

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/frdwin/trophy/cmd/apps"
)

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
	fuzzyFinder := exec.Command("/usr/bin/sk")
	fuzzyFinder.Stdin = strings.NewReader(strings.Join(appNames, "\n"))
	fuzzyFinder.Stdout = &stdout

	err = fuzzyFinder.Run()
	if err != nil {
		log.Fatalf("Error starting fuzzy finder: %s\n", err)
	}

	parsedStdout := strings.TrimRight(stdout.String(), "\n")
	app, err := appList.GetApp(parsedStdout)
	if err != nil {
		log.Fatalf("Error getting chosen app: %s\n", err)
	}

	execPath, err := exec.LookPath(strings.Trim(app.Cmd, " "))
	if err != nil {
		log.Fatalf("Error finding chosen app's path: %s\n", err)
	}

	env := os.Environ()
	args := []string{app.Cmd}
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
