package server

import (
	"github.com/denizakturk/servfront/kernel"
	"log"
	"net/http"
)

func registerRouter(w http.ResponseWriter, r *http.Request) {
	kernel.Holder.TokenCatcher(r)
	kernel.Holder.Checker()
	for _, routerHolder := range kernel.Holder.Config.Router.RouterHolder {
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
