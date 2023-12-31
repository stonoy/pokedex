package main

type command struct {
	name        string
	description string
	callback    func(*state, []string) error
}

func m() map[string]command {
	return map[string]command{
		"help": {
			name:        "Help",
			description: "Here to help u...",
			callback:    helpCallback,
		},
		"exit": {
			name:        "Exit",
			description: "Exit from current session",
			callback:    exitCallback,
		},
		"map": {
			name:        "Map",
			description: "Forward Fetch location from api",
			callback:    getNextLocations,
		},
		"mapb": {
			name:        "Map Back",
			description: "Backword Fetch location from api",
			callback:    getPreviousLocations,
		},
		"explore": {
			name:        "Get Pokemons",
			description: "explore pokemons in a given area",
			callback:    getPokemons,
		},
		"catch": {
			name:        "Catch Pokemon",
			description: "catch your fav pokemon in a area of your choice",
			callback:    catchPokemon,
		},
		"inspect": {
			name:        "Inspect Pokemon",
			description: "Inspect pokemons in your pokedex",
			callback:    inspectPokemons,
		},
		"pokedex": {
			name:        "Check Pokemons",
			description: "check available pokemons in your pokedex",
			callback:    pokedexx,
		},
	}
}
