package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/maximedelporte/pokegolang/internal"
)

func main() {
	commands := internal.GetCommands()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if cmd, ok := commands[scanner.Text()]; ok {
			cmd.Callback()
		} else {
			fmt.Println("Command doesn't exist.\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
