package models

type Starships struct {
	Name_St                string  `json:"name"`
	Model                  string  `json:"model"`
	Manufacturer           string  `json:"manufacturer"`
	Cost_in_credits        string  `json:"cost_in_credits"`
	Length                 string  `json:"length"`
	Max_atmosphering_speed string  `json:"max_atmosphering_speed"`
	Crew                   string  `json:"crew"`
	Passengers             string  `json:"passengers"`
	Cargo_capacity         string  `json:"cargo_capacity"`
	Consumables            string  `json:"consumables"`
	Hyperdrive_rating      string  `json:"hyperdrive_rating"`
	Starship_class         string  `json:"starship_class"`
	Film                   []Films `json:"films"`
}

type StarshipsResponse struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Results  []Starships `json:"results"`
}
