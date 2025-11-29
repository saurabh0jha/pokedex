package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (cl *Client) PokemonDetails(name string) (RespShallowPokemonDetails, error) {
	url := baseURL + fmt.Sprintf("/pokemon/%s/", name)

	if data, isCached := cl.cache.Get(url); isCached {
		pokemonResp := RespShallowPokemonDetails{}
		err := json.Unmarshal(data, &pokemonResp)
		if err != nil {
			return RespShallowPokemonDetails{}, err
		}

		return pokemonResp, nil
	}

	resp, err := cl.httpClient.Get(url)
	if err != nil {
		return RespShallowPokemonDetails{}, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowPokemonDetails{}, err
	}

	pokemonResp := RespShallowPokemonDetails{}
	err = json.Unmarshal(respData, &pokemonResp)
	if err != nil {
		return RespShallowPokemonDetails{}, err
	}

	cl.cache.Add(url, respData)

	return pokemonResp, nil

}
