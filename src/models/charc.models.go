package models

// Faire une structure pour le charactere
type Charc struct {
	Name_c      string      `json:"name"`
	Height      int         `json:"height"`
	Mass        int         `json:"mass"`
	Haire_color string      `json:"haire_color"`
	Birth_year  string      `json:"birth_year"`
	Gender      string      `json:"gender"`
	Species     []Species   `json:"species"`
	Homeworld   []Planets   `json:"homeworld"`
	Films       []Films     `json:"films"`
	Vehicle     []Veheicles `json:"vehicles"`
	Starship    []Starships `json:"starships"`
}
