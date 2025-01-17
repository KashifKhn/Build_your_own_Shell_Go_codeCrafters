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
		commad := extractCommand(reader)
		listOfCommands := []string{"exit", "echo", "type"}

		switch {
		case strings.HasPrefix(commad, "exit"):
			os.Exit(0)
		case strings.HasPrefix(commad, "echo"):
			echoText := strings.TrimSpace(strings.Split(commad, "echo")[1])
			fmt.Println(echoText)
		case strings.HasPrefix(commad, "type "):
			typeText := strings.TrimSpace(strings.Split(commad, "type ")[1])
			if slices.Contains(listOfCommands, typeText) {
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
			commandNotFound(commad)
		}
	}
}

func findAllExcutableCmd(args string) (string, bool) {
	paths := os.Getenv("PATH")
	pathsList := strings.Split(paths, ":")
	for _, path := range pathsList {
		dirs, _ := os.ReadDir(path)
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

func commandNotFound(command string) {
	fmt.Printf("%s: command not found\n", command)
}

func extractCommand(reader string) string {
	return strings.Split(strings.ToLower(strings.TrimSpace(reader)), "\n")[0]
}
