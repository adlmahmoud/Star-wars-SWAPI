package routers

import (
	"net/http"
	"swapi/src/controllers"
)

func FilmRouters(routers *http.ServeMux) {
	routers.HandleFunc("/films", controllers.Displayfilms)
}
