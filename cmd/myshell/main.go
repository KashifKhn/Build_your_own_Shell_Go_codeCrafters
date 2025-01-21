package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	for {
		fmt.Fprint(os.Stdout, "$ ")
		reader := userInput()
		cmd := extractCommand(reader)
		listOfCmds := map[string]bool{"exit": true, "echo": true, "type": true, "pwd": true, "cd": true}

		switch {
		case strings.HasPrefix(cmd, "exit"):
			os.Exit(0)
		case strings.HasPrefix(cmd, "echo"):
			echo(cmd)
		case strings.HasPrefix(cmd, "pwd"):
			pwd()
		case strings.HasPrefix(cmd, "cd"):
			cd(cmd)
		case strings.HasPrefix(cmd, "type"):
			args := strings.Fields(cmd)
			if len(args) < 2 {
				fmt.Println("type: missing argument")
				continue
			}
			typeCmd(args, listOfCmds)
		default:
			cmdFields := strings.Fields(cmd)
			cmdExec := cmdFields[0]
			agrs := cmdFields[1:]
			_, isFound := findAllExcutableCmd(cmdExec)
			if isFound {
				excuteCmd(cmdExec, agrs)
			} else {
				commandNotFound(cmd)
			}
		}
	}
}

func cd(cmd string) {
	cmdFields := strings.Fields(cmd)
	cmdExec := cmdFields[0]
	path := cmdFields[1]
	err := os.Chdir(path)
	if err != nil {
		fmt.Printf("%s: %s: No such file or directory\n", cmdExec, path)
	}
}

func typeCmd(args []string, listOfCmds map[string]bool) {
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
}

func pwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(dir)
}

func echo(cmd string) {
	args := strings.Fields(cmd)
	echoStr := strings.Join(args[1:], " ")
	fmt.Println(echoStr)
}

func excuteCmd(cmd string, args []string) {
	osCmd := exec.Command(cmd, args...)
	osCmd.Stdout = os.Stdout
	osCmd.Stderr = os.Stderr
	err := osCmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
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
