package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/maximedelporte/pokegolang/internal"
)

func main() {
	commands := internal.GetCommands()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		args := strings.Split(scanner.Text(), " ")
		stringCmd := args[0]

		if cmd, ok := commands[stringCmd]; ok {
			if stringCmd == "explore" || stringCmd == "catch" {
				if len(args) != 2 {
					fmt.Printf("command is invalid. Call help to see the usage.\n\n")
					return
				}
				argument := args[1]
				cmd.Callback(argument)
			} else {
				cmd.Callback(nil)
			}
		} else {
			fmt.Printf("Command doesn't exist.\n\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
