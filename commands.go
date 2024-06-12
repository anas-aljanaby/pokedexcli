package main 

import (
    "fmt"
    "os"
)

type cliCommand struct {
    name	string
    description	string
    callback	func() error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
	"help": {
	    name:	"help",
	    description:	"Displays a help message",
	    callback: commandHelp,
	},
	"exit": {
	    name:        "exit",
	    description: "Exit the Pokedex",
	    callback:    commandExit,
	},
	"map": {
	    name:        "map",
	    description: "Print the next 20 areas",
	    callback:   mapCommand,
	},
	"mapb": {
	    name:        "mapb",
	    description: "Print the previous 20 areas",
	    callback:   mapbCommand,
	},
    }
}

func commandHelp() error {
    commands := getCommands()
    for _, cmd := range commands {
	fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    fmt.Println()
    return nil
}

func commandExit() error {
    os.Exit(0)
    return nil
}

func printAreas(respB *respBatch) {
    for _, area := range respB.Results {
	fmt.Println(area.Name)
    }
}

func mapCommand() error {
    getBatch(&curResp, "n")
    printAreas(&curResp)
    return nil
}

func mapbCommand() error {
    getBatch(&curResp, "b")
    printAreas(&curResp)
    return nil
}
