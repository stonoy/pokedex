package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func read(newState *state) {
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print("pokedex -> ")
		scanner.Scan()

		inputs := scanner.Text()
		cleanedInput := clearText(inputs)

		if len(cleanedInput) > 0 {
			primaryCommand := cleanedInput[0]
			otherCommands := cleanedInput[1:]
			evaluate(primaryCommand, newState, otherCommands)
		}
	}
}

func clearText(s string) []string {
	lowerS := strings.ToLower(s)
	words := strings.Fields((lowerS))
	return words
}

func evaluate(str string, newState *state, other []string) {
	cmdToAction := m()
	value, ok := cmdToAction[str]
	if !ok {
		fmt.Println("Invalid command!!!")
		return
	}
	err := value.callback(newState, other)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func helpCallback(newState *state, other []string) error {
	if len(other) > 0 {
		fmt.Printf("extra command given, No need of %v and so on\n", other[0])
	}
	fmt.Println("Here are some commands...")
	cmdToAction := m()
	for key, _ := range cmdToAction {
		fmt.Printf("  ->> %v\n", key)
	}
	return nil
}

func exitCallback(newState *state, other []string) error {
	if len(other) > 0 {
		fmt.Printf("extra command given, No need of %v and so on\n", other[0])
	}
	os.Exit(0)
	return nil
}

func getNextLocations(newState *state, other []string) error {
	if len(other) > 0 {
		fmt.Printf("extra command given, No need of %v and so on\n", other[0])
	}
	locations, err := newState.newClient.MakeRequest(newState.nextUrl)
	if err != nil {
		return err
	}
	newState.editState(locations.Next, locations.Previous)
	for i, location := range locations.Results {
		fmt.Printf("%v -> %v\n", i, location.Name)
	}
	return nil
}

func getPreviousLocations(newState *state, other []string) error {
	if len(other) > 0 {
		fmt.Printf("extra command given, No need of %v and so on\n", other[0])
	}
	if newState.previousUrl == nil {
		return errors.New("You are at first page!")
	}

	locations, err := newState.newClient.MakeRequest(newState.previousUrl)
	if err != nil {
		return err
	}
	// fmt.Println(locations)
	newState.editState(locations.Next, locations.Previous)
	for i, location := range locations.Results {
		fmt.Printf("%v -> %v\n", i, location.Name)
	}
	return nil
}

func getPokemons(newState *state, other []string) error {
	if len(other) != 1 {
		return errors.New("problem in second argument")
	}

	pokemons, err := newState.newClient.AreaPokemons(other[0])
	if err != nil {
		return err
	}
	fmt.Printf("Available pokemons in location %v\n", other[0])
	for i, pokemon := range pokemons.PokemonEncounters {
		fmt.Printf("%v -> %v\n", i, pokemon.Pokemon.Name)
	}
	return nil
}

func catchPokemon(newState *state, other []string) error {
	if len(other) != 1 {
		return errors.New("problem in second argument")
	}

	fmt.Printf("throwing pokeball to %v\n", other[0])

	pokemon, err := newState.newClient.SinglePokemon(other[0])
	if err != nil {
		return err
	}

	threshold := 50
	strengthOfPokemon := rand.Intn(pokemon.BaseExperience)
	if threshold < strengthOfPokemon {
		fmt.Printf("%v escaped\n", other[0])
		return nil
	}

	fmt.Printf("%v was caugth\n", other[0])
	newState.addPokemonToPokedex(other[0], pokemon)

	return nil
}

func inspectPokemons(newState *state, other []string) error {
	pokemon, ok := newState.getPokemonFromPokedex(other[0])
	if !ok {
		fmt.Printf("you have not caught %v\n", other[0])
		return nil
	}

	fmt.Printf("Name: %v\n", other[0])
	fmt.Printf("Height %v\n", pokemon.Height)
	fmt.Printf("Weigth %v\n", pokemon.Weight)
	fmt.Println("Stats: ")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types: ")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%v\n", t.Type.Name)
	}

	return nil

}

func pokedexx(newState *state, other []string) error {
	for key, _ := range newState.pokedex {
		fmt.Printf("  - %v\n", key)
	}
	return nil
}
