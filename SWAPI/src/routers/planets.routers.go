package routers

import (
	"net/http"
	"swapi/src/controllers"
)

func PlanetsRouters(routers *http.ServeMux) {
	routers.HandleFunc("/planets", controllers.DisplayPlanets)
}
