package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var page int = 0

// var loc_cache pokecache.Cache = pokecache.NewCache(5*time.Second, time.Minute*5)

func getLocations(url string) (Locations, error) {
	// _, found := loc_cache.Get("locations")
	// if va {
	// 	location := Locations{}
	// 	err := json.Unmarshal(loc_cache["locations"].val, &location)
	// 	return location, err
	// }
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(baseUrl + url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	// pokecache.cache.Add(url, body)
	location := Locations{}
	err = json.Unmarshal(body, &location)
	return location, err
}
