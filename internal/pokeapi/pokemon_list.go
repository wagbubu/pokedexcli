package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
)

func (c *Client) PokemonList(mapName string) (EncounterData, error) {
	if mapName == "" {
		return EncounterData{}, errors.New("map name required in command explore <map-name>")
	}
	url := baseURL + "/location-area/" + mapName
	res, err := c.httpClient.Get(url)
	if err != nil {
		return EncounterData{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return EncounterData{}, err
	}

	pokemonResp := EncounterData{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return EncounterData{}, err
	}

	return pokemonResp, nil
}