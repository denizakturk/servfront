# Service Frontend

We did cluster all pieces on in kernel for stability and to be useful and starting servfront.

**installation:**
`go install github.com/denizakturk/servfront`

**Global Variable**
`kernel.Holder`

To simply serve content:
===
Create server variables
___
```go
	kernel.Holder.Init()
```
Create route
___
```go
	serviceController := &controller.Index{}

	routeName := "index_cluster"

	routePath := &router.RouteAddress{Pattern: "/hello/{name}"}

	route := &router.Route{
		Controller: serviceController,
		Endpoint:   serviceController.Hello,
		Name:       routeName,
		Address:    routePath,
	}
```
Register route to server
___
```go
	kernel.Holder.AddRoute("index_cluster", route)
```
Reading and prepare in kernel route and config parameters before start server
___
```go
	server.InitServer()
```

**Examples**
---

Example controller
___

```go

import (
	"net/http"
	"github.com/denizakturk/servfront/bridge"
)

type Index struct {
    ControllerMiddleware
}

func (c *Index) Hello()*bridge.ServiceResponse{
	return &bridge.ServiceResponse{}
}

```

Reusable Middleware
___

```go

type ControllerModdleware struct {
	response    http.ResponseWriter
    request     *http.Request
    urlParams   *bridge.UrlParams
}


func (cm *ControllerModdleware) MiddleWare(w http.ResponseWriter, r *http.Request, urlParams *bridge.UrlParams) {
    // Middleware codes for between service and controller
    cm.response = w
    cm.request = r
    cm.urlParams = urlParams
}

func (c *ControllerModdleware) EndpointRunner(method func() *bridge.ServiceResponse) {
    // Controller method run and write response
    response := method()
    response.WriteResponse(cm.response, nil)
}

```