package models

type Veheicles struct {
	Name_v                 string  `json:"name"`
	Model                  string  `json:"model"`
	Manufacturer           string  `json:"manufacturer"`
	Cost_in_credits        string  `json:"cost_in_credits"`
	Length                 string  `json:"length"`
	Max_atmosphering_speed string  `json:"max_atmosphering_speed"`
	Crew                   string  `json:"crew"`
	Passengers             string  `json:"passengers"`
	Cargo_capacity         string  `json:"cargo_capacity"`
	Consumables            string  `json:"consumables"`
	Vehicle_class          string  `json:"vehicle_class"`
	Film                   []Films `json:"films"`
}

type VeheiclesRep struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Results  []Veheicles `json:"results"`
}
