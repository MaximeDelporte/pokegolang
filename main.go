package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands := map[string]command{
		"help": {
			name: 		 "help",
			description: "Displays a help message",
			callback: 	 commandHelp,
		},
		"exit": {
			name: 		 "exit",
			description: "Exit the Pokedex",
			callback: 	 commandExit,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if cmd, ok := commands[scanner.Text()]; ok {
			cmd.callback()
		} else {
			fmt.Println("Command doesn't exist.\n")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

type command struct {
	name string
	description string
	callback func() error
}

func commandHelp() error {
	fmt.Printf("\nWelcome to the Pokedex!\n")
	fmt.Printf("Usage:\n\n")

	commands := map[string]command{
		"help": {
			name: 		 "help",
			description: "Displays a help message",
			callback: 	 commandHelp,
		},
		"exit": {
			name: 		 "exit",
			description: "Exit the Pokedex",
			callback: 	 commandExit,
		},
	}

	for _, command := range commands {
		fmt.Printf("%s: %s\n", command.name, command.description)
	}

	fmt.Println("")

	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}
