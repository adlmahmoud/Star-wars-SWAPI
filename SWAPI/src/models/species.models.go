package models

type Species struct {
	Name_s           string    `json:"name"`
	Classification   string    `json:"classification"`
	Designation      string    `json:"designation"`
	Average_height   int       `json:"average_height"`
	Skin_colors      string    `json:"skin_colors"`
	Hair_colors      string    `json:"hair_colors"`
	Eye_colors       string    `json:"eye_colors"`
	Average_lifespan int       `json:"average_lifespan"`
	Homeworld        []Planets `json:"homeworld"`
	Language         string    `json:"language"`
	People           []Charc   `json:"people"`
	Film             []Films   `json:"films"`
}

type SpeciesRep struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Species `json:"results"`
}
