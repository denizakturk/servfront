package config

import (
	"errors"
	"github.com/denizakturk/servfront/router"
)


type RouterHolder struct {
	Routes map[string]*router.Route
}

func (rh *RouterHolder) AddRouter(route *router.Route) error {
	if nil == rh.Routes {
		rh.Routes = make(map[string]*router.Route)
	}
	if _, ok := rh.Routes[route.Name]; !ok {
		rh.Routes[route.Name] = route
	} else {
		return errors.New("Duplicate router name | " + route.Name)
	}

	return nil
}

type RouterHolderCluster struct {
	RouterHolder map[string]*RouterHolder
}

func (rhc *RouterHolderCluster) AddRouterToCluster(clusterName string, route *router.Route) {
	if nil == rhc.RouterHolder {
		rhc.RouterHolder = make(map[string]*RouterHolder)
	}
	if _, ok := rhc.RouterHolder[clusterName]; !ok {
		routerHolder := &RouterHolder{}
		rhc.RouterHolder[clusterName] = routerHolder
	}
	rhc.RouterHolder[clusterName].AddRouter(route)
}

var Routes = &RouterHolderCluster{}
