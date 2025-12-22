package models

type Films struct {
	Titre         string      `json:"title"`
	Opening_crawl string      `json:"opening_crawl"`
	Director      string      `json:"director"`
	Producer      string      `json:"producer"`
	Release_date  int         `json:"release_date"`
	Characters    []Charc     `json:"characters"`
	Planet        []Planets   `json:"planets"`
	Starship      []Starships `json:"starships"`
	Vehicle       []Veheicles `json:"vehicles"`
	Specie        []Species   `json:"species"`
}
