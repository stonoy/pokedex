package ApiCalls

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://pokeapi.co/api/v2"

func (c *Client) MakeRequest(url *string) (LocationAreas, error) {
	endPoint := "/location-area"
	finalUrl := baseUrl + endPoint

	if url != nil {
		finalUrl = *url
	}

	data, ok := c.cacheMemory.GetArea(finalUrl)
	if ok {
		fmt.Println("area cache hit...")
		locations := LocationAreas{}
		err := json.Unmarshal(data, &locations)
		if err != nil {
			return LocationAreas{}, err
		}

		return locations, nil
	}

	r1, err := http.NewRequest("GET", finalUrl, nil)
	if err != nil {
		return LocationAreas{}, err
	}

	r2, err := c.httpClient.Do(r1)
	if err != nil {
		return LocationAreas{}, err
	}

	defer r2.Body.Close()

	d1, err := io.ReadAll(r2.Body)
	if err != nil {
		return LocationAreas{}, err
	}

	locations := LocationAreas{}
	err = json.Unmarshal(d1, &locations)
	if err != nil {
		return LocationAreas{}, err
	}

	c.cacheMemory.AddArea(finalUrl, d1)

	return locations, nil

}

func (c *Client) AreaPokemons(areaName string) (Pokemons, error) {
	endPoint := "/location-area/" + areaName
	finalUrl := baseUrl + endPoint

	pokemonBytes, ok := c.cacheMemory.GetPokemons(areaName)
	if ok {
		fmt.Println("pokemon caches hits..")
		pokemons := Pokemons{}

		err := json.Unmarshal(pokemonBytes, &pokemons)
		if err != nil {
			return Pokemons{}, err
		}
		return pokemons, nil
	}

	r1, err := http.NewRequest("GET", finalUrl, nil)
	if err != nil {
		return Pokemons{}, err
	}

	r2, err := c.httpClient.Do(r1)
	if err != nil {
		return Pokemons{}, err
	}

	defer r2.Body.Close()

	d1, err := io.ReadAll(r2.Body)
	if err != nil {
		return Pokemons{}, err
	}

	pokemons := Pokemons{}

	err = json.Unmarshal(d1, &pokemons)
	if err != nil {
		return Pokemons{}, err
	}

	c.cacheMemory.AddPokemons(areaName, d1)

	return pokemons, nil

}

func (c *Client) SinglePokemon(name string) (Pokemon, error) {
	endPoint := "/pokemon/" + name
	finalUrl := baseUrl + endPoint

	r1, err := http.NewRequest("GET", finalUrl, nil)
	if err != nil {
		return Pokemon{}, err
	}

	r2, err := c.httpClient.Do(r1)
	if err != nil {
		return Pokemon{}, err
	}

	defer r2.Body.Close()

	d1, err := io.ReadAll(r2.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}

	err = json.Unmarshal(d1, &pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil

}
