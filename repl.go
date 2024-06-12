package main

import (
    "fmt"
    "strings"
    "bufio"
    "os"
)

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)
    for {
	fmt.Print("Pokedex > ")
	if !scanner.Scan() {
	    break
	}
	input := scanner.Text()
	words := cleanInput(input)
	if len(words) == 0 {
	    continue
	}
	cmdName := words[0]
	args := []string{}
	if len(words) > 1 {
	    args = words[1:]
	}
	cmd, exits := getCommands()[cmdName]
	if exits {
	    err := cmd.callback(args...)
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
