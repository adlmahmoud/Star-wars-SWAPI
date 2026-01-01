package routers

import (
	"net/http"
	"swapi/src/controllers"
)

func SpeciesRouters(reouters *http.ServeMux) {
	reouters.HandleFunc("/species", controllers.DisplaySpecies)
}
