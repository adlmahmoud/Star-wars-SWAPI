package routers

import (
	"net/http"
)

func MainRouter() *http.ServeMux {
	mainRouter := http.NewServeMux()

	errorRouter(mainRouter)
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
