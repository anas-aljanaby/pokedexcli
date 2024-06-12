package main

import (
    "time"
    "github.com/anas-aljanaby/pokedexcli/internal/pokecache"
)

var curResp = respBatch{
    Next: "https://pokeapi.co/api/v2/location-area?offset=0&limit=20",
}

var cache *pokecache.Cache

func main() {
    cache = pokecache.NewCache(5 * time.Minute)
    startRepl()
}
