package main

import (
	"fmt"
	"time"

	"github.com/stonoy/pokedex/Internal/ApiCalls"
)

type state struct {
	newClient   *ApiCalls.Client
	pokedex     map[string]ApiCalls.Pokemon
	nextUrl     *string
	previousUrl *string
}

func (s *state) editState(s1, s2 *string) {

	s.nextUrl = s1
	s.previousUrl = s2
}

func (s *state) addPokemonToPokedex(name string, pokemon ApiCalls.Pokemon) {
	s.pokedex[name] = pokemon
	fmt.Printf("%v added to my pokedex\n", name)
}

func (s *state) getPokemonFromPokedex(name string) (ApiCalls.Pokemon, bool) {
	val, ok := s.pokedex[name]
	if !ok {
		return ApiCalls.Pokemon{}, false
	}

	return val, true
}

func main() {
	fmt.Println("Welcome")
	NewState := &state{newClient: ApiCalls.NewClient(time.Minute, time.Hour), pokedex: map[string]ApiCalls.Pokemon{}}
	go NewState.newClient.RemoveCache()
	read(NewState)

}
