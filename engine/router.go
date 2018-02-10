package engine

import (
	"github.com/labstack/echo"
)

/*

 * This file contains the routing wrapper around the echo handler and routing

 TODO: Refactor the way current Verbs are handled and maybe do it dynamically.
 This will work for now and we can just expand it to the other verbs but might
 not be the best way to do it

*/

type RouteGroup struct {
	verbs    []string
	routes   []string
	handlers []echo.HandlerFunc
}

func (r *RouteGroup) Get(route string, hf echo.HandlerFunc) {
	r.addRoute("GET", route, hf)
}

func (r *RouteGroup) Post(route string, hf echo.HandlerFunc) {
	r.addRoute("POST", route, hf)
}

func (r *RouteGroup) Delete(route string, hf echo.HandlerFunc) {
	r.addRoute("DELETE", route, hf)
}

func (r *RouteGroup) addRoute(verb, route string, hf echo.HandlerFunc) {
	r.verbs = append(r.verbs, verb)
	r.routes = append(r.routes, route)
	r.handlers = append(r.handlers, hf)
}
