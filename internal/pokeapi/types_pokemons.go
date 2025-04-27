package pokeapi

type EncounterData struct {
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"` 
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
	URL string `json:"url"`
	BaseExperience int `json:"base_experience"`
	Height int `json:"height"`
	Weight int `json:"weight"`
	Stats []Stat `json:"stats"`
	Types []Types `json:"types"`
}

type Stat struct {
	BaseStat int `json:"base_stat"`
	Stat StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
}


type Types struct {
	Type TypeInfo `json:"type"`
}

type TypeInfo struct {
	Name string `json:"name"`
}