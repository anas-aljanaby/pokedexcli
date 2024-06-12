package main

import (
    "fmt"
    "strings"
)

func startRepl() {
    var input string
    for {
	fmt.Print("Pokedex > ")
	fmt.Scanln(&input)
	words := cleanInput(input)
	if len(words) == 0 {
	    continue
	}
	cmdName := words[0]
	cmd, exits := getCommands()[cmdName]
	if exits {
	    err := cmd.callback()
	    if err != nil {
		fmt.Println(err)
	    }
	} else {
	    fmt.Println("Unknown command")
	}
    }
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}
