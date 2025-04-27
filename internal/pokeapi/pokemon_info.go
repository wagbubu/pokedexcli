package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		errMsg := fmt.Sprintf("pokemon not found! are you sure %q exists?", pokemonName)
		return Pokemon{}, errors.New(errMsg)
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	json.Unmarshal(data, &pokemon)

	// prettyJSON, err := json.MarshalIndent(pokemon, "", "  ") // 2 spaces indent
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println(string(prettyJSON))


	return pokemon, nil
}