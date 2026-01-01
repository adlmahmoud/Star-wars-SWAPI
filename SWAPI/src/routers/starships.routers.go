package routers

import (
	"net/http"
	"swapi/src/controllers"
)

func StarshipsRouters(routers *http.ServeMux) {
	routers.HandleFunc("/starships", controllers.DisplayStarships)
}
