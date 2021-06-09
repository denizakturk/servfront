package config

import (
	"errors"
	"regexp"
	"github.com/denizakturk/servfront/bridge"
)

type Router struct {
	Name       string
	Pattern    *regexp.Regexp
	Controller bridge.Struct
	Endpoint   func() *bridge.ServiceResponse
}

type RouterHolder struct {
	Routers map[string]*Router
}

func (rh *RouterHolder) AddRouter(router *Router) error {
	if nil == rh.Routers {
		rh.Routers = make(map[string]*Router)
	}
	if _, ok := rh.Routers[router.Name]; !ok {
		rh.Routers[router.Name] = router
	} else {
		return errors.New("Duplicate router name | " + router.Name)
	}

	return nil
}

type RouterHolderCluster struct {
	RouterHolder map[string]*RouterHolder
}

func (rhc *RouterHolderCluster) AddRouterToCluster(clusterName string, router *Router) {
	if nil == rhc.RouterHolder {
		rhc.RouterHolder = make(map[string]*RouterHolder)
	}
	if _, ok := rhc.RouterHolder[clusterName]; !ok {
		routerHolder := &RouterHolder{}
		rhc.RouterHolder[clusterName] = routerHolder
	}
	rhc.RouterHolder[clusterName].AddRouter(router)
}

var Routes = &RouterHolderCluster{}
