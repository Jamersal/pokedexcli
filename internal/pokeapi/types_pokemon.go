package pokeapi

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
}
type PokemonStat struct {
	Name string `json:"name"`
}

type Stat struct {
	BaseStat int         `json:"base_stat"`
	Stat     PokemonStat `json:"stat"`
}

type PokemonType struct {
	Name string `json:"name"`
}
type Type struct {
	Type PokemonType `json:"type"`
}
