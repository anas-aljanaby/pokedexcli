package main 

import (
    "fmt"
    "os"
    "errors"
    "math/rand"
)

type cliCommand struct {
    name	string
    description	string
    callback	func(...string) error
}

func getCommands() map[string]cliCommand {
    return map[string]cliCommand{
	"help": {
	    name:	"help", description:	"Displays a help message",
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
	"catch": {
	    name:        "catch",
	    description: "catch a pokemon",
	    callback:   catchCommand,
	},
	"inspect": {
	    name:        "inspect",
	    description: "inspect a pokemon",
	    callback:   inspectCommand,
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
	return errors.New("explore command expects only one arg (location name)")
    }
    getLocation(&locations, args[0])
    for _, pokemon := range locations.PokemonEncounters  {
	fmt.Println(pokemon.Pokemon.Name)
    // fmt.Println(string(body))
    }
    return nil
}

func tryToCatch(pokemon Pokemon) bool {
    res := rand.Intn(pokemon.BaseExperience)
    return res < 40
}

func catchCommand(args ...string) error {
    if len(args) != 1 {
	return errors.New("catch command expects only one arg (pokemon name)")
    }
    pokemonName := args[0]
    err := getPokemon(&pokemon, args[0])
    if err != nil {
	fmt.Println(err)
	return nil
    }
    fmt.Printf("Throwing a Pokeball at %s..\n", pokemonName)
    _, ok := pokedex[pokemonName]
    if ok {
	fmt.Printf("%s already in pokedex\n", pokemonName)
	return nil
    }
    if tryToCatch(pokemon) {
	fmt.Printf("%s was caught!\n", pokemonName)
	pokedex[pokemonName] = pokemon
    } else {
	fmt.Printf("%s escaped!\n", pokemonName)
    }
    return nil
}

func inspectCommand(args ...string) error {
    if len(args) != 1 {
	return errors.New("inspect command expects only one arg (pokemon name)")
    }
    pokemonName := args[0]
    pokemon, ok := pokedex[pokemonName]
    if !ok {
	fmt.Printf("%s is not caught yet\n", pokemonName)
	return nil
    }
    fmt.Printf(
	"Name: %s\nHeight: %d\nWeight:%d\n",
	pokemon.Name, pokemon.Height, pokemon.Weight,
    )
    fmt.Println("Stats: ")
    for _, stat := range pokemon.Stats {
	fmt.Printf(" -%s: %d\n", stat.Stat.Name, stat.BaseStat)
    }
    fmt.Println("Types: ")
    for _, pokeType := range pokemon.Types {
	fmt.Printf(" -%s\n", pokeType.Type.Name)
    }
    return nil
}
