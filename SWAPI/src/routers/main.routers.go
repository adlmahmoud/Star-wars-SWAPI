package routers

import (
	"net/http"
	"swapi/src/controllers"
)

func MainRouter() *http.ServeMux {
	mainRouter := http.NewServeMux()
	mainRouter.HandleFunc("/", controllers.DisplayIndex)
	SpeciesRouters(mainRouter)
	CharcRouters(mainRouter)
	PlanetsRouters(mainRouter)
	StarshipsRouters(mainRouter)
	VehicilesRouters(mainRouter)
	FilmRouters(mainRouter)
	errorRouter(mainRouter)

	fileServer := http.FileServer(http.Dir("./../../assets"))
	mainRouter.Handle("/static/", http.StripPrefix("/static/", fileServer))

	return mainRouter
}
