package models

type Veheicles struct {
	Name_v                 string  `json:"name"`
	Model                  string  `json:"model"`
	Manufacturer           string  `json:"manufacturer"`
	Cost_in_credits        int     `json:"cost_in_credits"`
	Length                 float64 `json:"length"`
	Max_atmosphering_speed int     `json:"max_atmosphering_speed"`
	Crew                   int     `json:"crew"`
	Passengers             int     `json:"passengers"`
	Cargo_capacity         int     `json:"cargo_capacity"`
	Consumables            string  `json:"consumables"`
	Vehicle_class          string  `json:"vehicle_class"`
	Film                   []Films `json:"films"`
}
