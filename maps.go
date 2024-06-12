package main 

import (
    "net/http"
    "log"
    "io"
    "encoding/json"
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

    res, err := http.Get(url)
    if err != nil {
	log.Fatal(err)
    }
    body, err := io.ReadAll(res.Body)
    res.Body.Close()
    if res.StatusCode > 299 {
	    log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
    }
    if err != nil {
	    log.Fatal(err)
    }
    cache.Add(url, body)
    json.Unmarshal(body, &respB)
}

