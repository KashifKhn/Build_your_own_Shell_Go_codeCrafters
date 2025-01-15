package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		reader := userInput()
		commad := extractCommand(reader)

		switch {
		case strings.HasPrefix(commad, "exit"):
			os.Exit(0)
		case strings.HasPrefix(commad, "echo"):
			echoText := strings.TrimSpace(strings.Split(commad, "echo")[1])
			fmt.Println(echoText)
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
