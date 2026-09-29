package pokeapi

type LocationAreasResponse struct {
	Count    int     `json:"count"` //`xx` is a struct tag - extra metadata
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type LocationAreaDetail struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	GameIndex            int    `json:"game_index"`
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Names []struct {
		Name     string `json:"name"`
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
			MaxChance        int `json:"max_chance"`
			EncounterDetails []struct {
				MinLevel int `json:"min_level"`
				MaxLevel int `json:"max_level"`
				Chance   int `json:"chance"`
				Method   struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				ConditionValues []any `json:"condition_values"`
				PokemonDetails  any   `json:"pokemon_details"`
			} `json:"encounter_details"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

// {
//   "count": 1539,
//   "next": "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
//   "previous": null,
//   "results": [
//     {
//       "name": "canalave-city-area",
//       "url": "https://pokeapi.co/api/v2/location-area/1/"
//     },
//     {
//       "name": "eterna-city-area",
//       "url": "https://pokeapi.co/api/v2/location-area/2/"
//     },
//     {
//       "name": "pastoria-city-area",
//       "url": "https://pokeapi.co/api/v2/location-area/3/"
//     },
//     {
//       "name": "sunyshore-city-area",
//       "url": "https://pokeapi.co/api/v2/location-area/4/"
//     },
//     {
//       "name": "sinnoh-pokemon-league-area",
//       "url": "https://pokeapi.co/api/v2/location-area/5/"
//     },
//     {
//       "name": "oreburgh-mine-1f",
//       "url": "https://pokeapi.co/api/v2/location-area/6/"
//     },
//     {
//       "name": "oreburgh-mine-b1f",
//       "url": "https://pokeapi.co/api/v2/location-area/7/"
//     },
//     {
//       "name": "valley-windworks-area",
//       "url": "https://pokeapi.co/api/v2/location-area/8/"
//     },
//     {
//       "name": "eterna-forest-area",
//       "url": "https://pokeapi.co/api/v2/location-area/9/"
//     },
//     {
//       "name": "fuego-ironworks-area",
//       "url": "https://pokeapi.co/api/v2/location-area/10/"
//     },
//     {
//       "name": "mt-coronet-1f-route-207",
//       "url": "https://pokeapi.co/api/v2/location-area/11/"
//     },
//     {
//       "name": "mt-coronet-2f",
//       "url": "https://pokeapi.co/api/v2/location-area/12/"
//     },
//     {
//       "name": "mt-coronet-3f",
//       "url": "https://pokeapi.co/api/v2/location-area/13/"
//     },
//     {
//       "name": "mt-coronet-exterior-snowfall",
//       "url": "https://pokeapi.co/api/v2/location-area/14/"
//     },
//     {
//       "name": "mt-coronet-exterior-blizzard",
//       "url": "https://pokeapi.co/api/v2/location-area/15/"
//     },
//     {
//       "name": "mt-coronet-4f",
//       "url": "https://pokeapi.co/api/v2/location-area/16/"
//     },
//     {
//       "name": "mt-coronet-4f-small-room",
//       "url": "https://pokeapi.co/api/v2/location-area/17/"
//     },
//     {
//       "name": "mt-coronet-5f",
//       "url": "https://pokeapi.co/api/v2/location-area/18/"
//     },
//     {
//       "name": "mt-coronet-6f",
//       "url": "https://pokeapi.co/api/v2/location-area/19/"
//     },
//     {
//       "name": "mt-coronet-1f-from-exterior",
//       "url": "https://pokeapi.co/api/v2/location-area/20/"
//     }
//   ]
// }
