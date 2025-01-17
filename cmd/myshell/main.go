package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		reader := userInput()
		cmd := extractCommand(reader)
		listOfCmds := map[string]bool{"exit": true, "echo": true, "type": true}

		switch {
		case strings.HasPrefix(cmd, "exit"):
			os.Exit(0)
		case strings.HasPrefix(cmd, "echo"):
			args := strings.Fields(cmd)
			echoStr := strings.Join(args[1:], " ")
			fmt.Println(echoStr)
		case strings.HasPrefix(cmd, "type"):
			args := strings.Fields(cmd)
			if len(args) < 2 {
				fmt.Println("type: missing argument")
				continue
			}
			cmdName := args[1]
			if listOfCmds[cmdName] {
				fmt.Printf("%s is a shell builtin\n", cmdName)
			} else {
				path, isFound := findAllExcutableCmd(cmdName)
				if isFound {
					fmt.Printf("%s is %s\n", cmdName, path)
				} else {
					fmt.Printf("%s: not found\n", cmdName)
				}
			}
		default:
			commandNotFound(cmd)
		}
	}
}

func findAllExcutableCmd(args string) (string, bool) {
	paths := os.Getenv("PATH")
	if paths == "" {
		return "", false
	}
	pathsList := strings.Split(paths, ":")
	for _, path := range pathsList {
		dirs, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		for _, dir := range dirs {
			if dir.Name() == args {
				return filepath.Join(path, dir.Name()), true
			}
		}
	}
	return "", false
}

func userInput() string {
	reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading from stdin")
		os.Exit(1)
	}
	return strings.TrimSpace(reader)
}

func commandNotFound(cmd string) {
	fmt.Printf("%s: command not found\n", cmd)
}

func extractCommand(reader string) string {
	return strings.Split(strings.TrimSpace(reader), "\n")[0]
}
