package server

import (
	"errors"
	"log"
	"net/http"

	"github.com/denizakturk/servfront/bridge"
	"github.com/denizakturk/servfront/kernel"
	"github.com/denizakturk/servfront/router"
)

func registerRouter(w http.ResponseWriter, r *http.Request) {
	kernel.Holder.TokenCatcher(r)
	matchUrl := false
	var runnableRouter *router.Route = nil
BreakHolder:
	for _, routerHolder := range kernel.Holder.Config.Router.RouterHolder {
		for _, router := range routerHolder.Routes {
			if nil != router.Pattern && router.Pattern.MatchString(r.URL.Path) {
				runnableRouter = router
				runnableRouter.Controller.MiddleWare(w, r, &bridge.UrlParams{})
				matchUrl = true
				break BreakHolder
			} else if nil != router.Address && router.Address.RegexpPattern.MatchString(r.URL.Path) {
				runnableRouter = router
				runnableRouter.Address.CatchAddressParametersValue(r.URL.Path)
				runnableRouter.Controller.MiddleWare(w, r, &bridge.UrlParams{Params: router.Address.ParamsToMap()})
				matchUrl = true
				break BreakHolder
			}
		}
	}
	if matchUrl {
		if runnableRouter.TokenValidate == nil || *runnableRouter.TokenValidate {
			checkerError := kernel.Holder.Checker()
			if checkerError != nil {
				response := &bridge.ServiceResponse{Error: checkerError}
				response.WriteResponse(w)
				return
			}
		}

		runnableRouter.Controller.EndpointRunner(runnableRouter.Endpoint)
		return
	}
	if !matchUrl {
		response := &bridge.ServiceResponse{Error: errors.New("Url not found: " + r.URL.Path)}
		response.WriteResponse(w)
	}
}

func InitServer() {
	http.HandleFunc("/", registerRouter)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
