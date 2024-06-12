package main 

import (
    "fmt"
    "os"
    "errors"
)

type cliCommand struct {
    name	string
    description	string
    callback	func(...string) error
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
	"explore": {
	    name:        "explore",
	    description: "Explore an area",
	    callback:   exploreCommand,
	},
    }
}

func commandHelp(args ...string) error {
    commands := getCommands()
    for _, cmd := range commands {
	fmt.Printf("%s: %s\n", cmd.name, cmd.description)
    }
    fmt.Println()
    return nil
}

func commandExit(args ...string) error {
    os.Exit(0)
    return nil
}

func printAreas(respB *respBatch) {
    for _, area := range respB.Results {
	fmt.Println(area.Name)
    }
}

func mapCommand(args ...string) error {
    getBatch(&curResp, "n")
    printAreas(&curResp)
    return nil
}

func mapbCommand(args ...string) error {
    getBatch(&curResp, "b")
    printAreas(&curResp)
    return nil
}

func exploreCommand(args ...string) error {
    if len(args) != 1 {
	return errors.New("explore command expects arg (location name)")
    }
    getLocation(&locations, args[0])
    for _, pokemon := range locations.PokemonEncounters  {
	fmt.Println(pokemon.Pokemon.Name)
    // fmt.Println(string(body))
    }
    return nil
}
