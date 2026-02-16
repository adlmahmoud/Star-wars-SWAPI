package models

type Planets struct {
	Name_p          string  `json:"name"`
	Rotation_period string  `json:"rotation_period"`
	Orbital_period  string  `json:"orbital_period"`
	Diameter        string  `json:"diameter"`
	Climate         string  `json:"climate"`
	Gravity         string  `json:"gravity"`
	Terrain         string  `json:"terrain"`
	Population      string  `json:"population"`
	Residents       []Charc `json:"residents"`
	Films           []Films `json:"films"`
}

type PlanetsResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Planets `json:"results"`
}
