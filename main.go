// Author: James Mallon <jamesmallondev@gmail.com>
// Main package
package main

import (
	//"axeman/libs/logger"
	"axeman/routes"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// starting point...
func main() {
	router := httprouter.New()

	routes.GetRoutes(router)
	routes.GetAuthRoutes(router)

	router.ServeFiles("/static/*filepath", http.Dir("public/"))
	router.ServeFiles("/libs/*filepath", http.Dir("node_modules/"))
	panic(http.ListenAndServe("localhost:8080", router))
}
