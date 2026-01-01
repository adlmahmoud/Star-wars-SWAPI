package routers

import (
	"net/http"
	"swapi/src/controllers"
)

func CharcRouters(routers *http.ServeMux) {
	routers.HandleFunc("/characters", controllers.DisplayCharc)
}
