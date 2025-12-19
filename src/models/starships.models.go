package models

type Starships struct {
	Name_St                string  `json:"name"`
	Model                  string  `json:"model"`
	Manufacturer           string  `json:"manufacturer"`
	Cost_in_credits        int     `json:"cost_in_credits"`
	Length                 int     `json:"length"`
	Max_atmosphering_speed int     `json:"max_atmosphering_speed"`
	Crew                   int     `json:"crew"`
	Passengers             int     `json:"passengers"`
	Cargo_capacity         int     `json:"cargo_capacity"`
	Consumables            string  `json:"consumables"`
	Hyperdrive_rating      int     `json:"hyperdrive_rating"`
	Starship_class         string  `json:"starship_class"`
	Film                   []Films `json:"films"`
}
