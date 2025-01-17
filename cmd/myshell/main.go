package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		reader := userInput()
		cmd := extractCommand(reader)
		listOfCmd := []string{"exit", "echo", "type"}

		switch {
		case strings.HasPrefix(cmd, "exit"):
			os.Exit(0)
		case strings.HasPrefix(cmd, "echo"):
			echoText := strings.TrimSpace(strings.Split(cmd, "echo")[1])
			fmt.Println(echoText)
		case strings.HasPrefix(cmd, "type "):
			typeText := strings.TrimSpace(strings.Split(cmd, "type ")[1])
			if slices.Contains(listOfCmd, typeText) {
				fmt.Printf("%s is a shell builtin\n", typeText)
			} else {
				path, isFound := findAllExcutableCmd(typeText)
				if isFound {
					fmt.Printf("%s is %s\n", typeText, path)
				} else {
					fmt.Printf("%s: not found\n", typeText)
				}
			}
		default:
			commandNotFound(cmd)
		}
	}
}

func findAllExcutableCmd(args string) (string, bool) {
	paths := os.Getenv("PATH")
	pathsList := strings.Split(paths, ":")
	for _, path := range pathsList {
		dirs, err := os.ReadDir(path)
		if err != nil {
			continue
		}
		for _, dir := range dirs {
			if dir.Name() == args {
				str := path + "/" + dir.Name()
				return str, true
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
	return reader
}

func commandNotFound(cmd string) {
	fmt.Printf("%s: command not found\n", cmd)
}

func extractCommand(reader string) string {
	return strings.Split(strings.TrimSpace(reader), "\n")[0]
}
