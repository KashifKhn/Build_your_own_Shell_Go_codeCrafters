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
		reader, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from stdin")
			return
		}

		fmt.Printf("%s: command not found\n", strings.Split(reader, "\n")[0])

	}

}
