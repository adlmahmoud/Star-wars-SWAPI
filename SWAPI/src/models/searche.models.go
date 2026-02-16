package models

// bon vue que l'api n'as pas un endpoint de recherche globale je me suis dit, je vais faire une struct global qui stock les donnés
type GlobalSearchResult struct {
	Characters []Charc
	Planets    []Planets
	Films      []Films
	Species    []Species
	Vehicles   []Veheicles
	Starships  []Starships
	Query      string
	Count      int
}
