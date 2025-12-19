package models

type Planets struct {
	Name_p          string  `json:"name"`
	Rotation_period int     `json:"rotation_period"`
	Orbital_period  int     `json:"orbital_period"`
	Diameter        int     `json:"diameter"`
	Climate         string  `json:"climate"`
	Gravity         string  `json:"gravity"`
	Terrain         string  `json:"terrain"`
	Population      int     `json:"population"`
	Residents       []Charc `json:"residents"`
	Films           []Films `json:"films"`
}
