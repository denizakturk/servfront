package server

import (
	"log"
	"net/http"
	"github.com/denizakturk/servfront/config"
)

func registerRouter(w http.ResponseWriter, r *http.Request) {
	for _, routerHolder := range config.Routes.RouterHolder {
		for _, router := range routerHolder.Routers {
			if router.Pattern.MatchString(r.URL.Path) {
				router.Controller.MiddleWare(w, r)
				router.Controller.EndpointRunner(router.Endpoint)
			}
		}
	}
}

func InitServer() {
	http.HandleFunc("/", registerRouter)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
