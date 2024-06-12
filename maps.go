package main 

import (
    "net/http"
    "log"
    "io"
    "encoding/json"
    "errors"
    "fmt"
)


type respBatch struct {
    Count    int    `json:"count"`
    Next     string `json:"next"`
    Previous string    `json:"previous"`
    Results  []struct {
	Name string `json:"name"`
	URL  string `json:"url"`
    } `json:"results"`
}

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

func fetchData(url string) ([]byte, error) {
    res, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, err
    }
    if res.StatusCode > 299 {
	return nil, errors.New(
	    fmt.Sprintf(
		"Response failed with status code: %d and\nbody: %s\n",
		res.StatusCode, body,
	    ),
	)
    }
    return body, nil
}

func fetchDataPokemon(url string) ([]byte, error, int) {
    res, err := http.Get(url)
    if err != nil {
        return nil, err, 0
    }
    defer res.Body.Close()
    body, err := io.ReadAll(res.Body)
    if err != nil {
        return nil, err, 0
    }
    if res.StatusCode > 299 {
	return nil, nil, res.StatusCode
    }
    return body, nil, 0
}

func getBatch(respB *respBatch, batch string) {
    var url string
    if batch == "n" {
	url = respB.Next
    } else if batch == "b" {
	url = respB.Previous
    } else {
	url = batch
    }

    if cachedData, ok := cache.Get(url); ok {
	json.Unmarshal(cachedData, &respB)
	return
    }

    body, err := fetchData(url)
    if err != nil {
        log.Fatal(err)
    }

    cache.Add(url, body)
    json.Unmarshal(body, &respB)
}

func getLocation(respN *Location, locationName string) {
    url := "https://pokeapi.co/api/v2/location-area/" + locationName
    // fmt.Println(url)
    body, err := fetchData(url)
    if err != nil {
        log.Fatal(err)
    }
    json.Unmarshal(body, &respN)
}

func getPokemon(pokemon *Pokemon, pokemonName string) error {
    url := "https://pokeapi.co/api/v2/pokemon/" + pokemonName
    body, err, code := fetchDataPokemon(url)
    if err != nil {
        log.Fatal(err)
    }
    if code == 404 {
	return errors.New("Not valid pokemon name")
    }
    json.Unmarshal(body, &pokemon)
    return nil
}
