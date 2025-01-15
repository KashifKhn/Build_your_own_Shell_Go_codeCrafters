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
				fmt.Printf("%s: not found\n", typeText)
			}
		default:
			commandNotFound(commad)
		}
	}
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
