package server

import (
	"errors"
	"github.com/denizakturk/servfront/bridge"
	"github.com/denizakturk/servfront/kernel"
	"log"
	"net/http"
)

func registerRouter(w http.ResponseWriter, r *http.Request) {
	kernel.Holder.TokenCatcher(r)
	checkerError := kernel.Holder.Checker()
	matchUrl := false
	if checkerError != nil {
		response := &bridge.ServiceResponse{Error: checkerError}
		response.WriteResponse(w, nil)
		return
	}
BreakHolder:
	for _, routerHolder := range kernel.Holder.Config.Router.RouterHolder {
		for _, router := range routerHolder.Routes {
			if nil != router.Pattern && router.Pattern.MatchString(r.URL.Path) {
				router.Controller.MiddleWare(w, r, &bridge.UrlParams{})
				router.Controller.EndpointRunner(router.Endpoint)
				matchUrl = true
				break BreakHolder
			} else if nil != router.Address && router.Address.RegexpPattern.MatchString(r.URL.Path) {
				router.Address.CatchAddressParametersValue(r.URL.Path)
				router.Controller.MiddleWare(w, r, &bridge.UrlParams{Params: router.Address.ParamsToMap()})
				router.Controller.EndpointRunner(router.Endpoint)
				matchUrl = true
				break BreakHolder
			}
		}
	}
	if !matchUrl {
		response := &bridge.ServiceResponse{Error: errors.New("Url not found: " + r.URL.Path)}
		errorCode := 404
		response.WriteResponse(w, &errorCode)
	}
}

func InitServer() {
	http.HandleFunc("/", registerRouter)
	log.Fatal(http.ListenAndServe(":9000", nil))
}
