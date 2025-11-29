package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

func (cl *Client) ExploreLocation(name string) (RespShallowLocationDetails, error) {
	url := baseURL + fmt.Sprintf("/location-area/%s", name)

	if data, isCached := cl.cache.Get(url); isCached {
		exploreResp := RespShallowLocationDetails{}
		fmt.Printf("Exploration data for %s from cache \n", url)
		err := json.Unmarshal(data, &exploreResp)
		if err != nil {
			return RespShallowLocationDetails{}, err
		}

		return exploreResp, nil
	}

	fmt.Printf("Data of exploration for %s not found in cache \n", url)

	resp, err := cl.httpClient.Get(url)
	if err != nil {
		return RespShallowLocationDetails{}, err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocationDetails{}, err
	}

	exploreResp := RespShallowLocationDetails{}
	err = json.Unmarshal(respData, &exploreResp)
	if err != nil {
		return RespShallowLocationDetails{}, err
	}

	cl.cache.Add(url, respData)

	return exploreResp, nil

}
