// Author: James Mallon <jamesmallondev@gmail.com>
// Main package
package main

import (
	"axeman/libs/databases"
	"axeman/models"
	"axeman/routes"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// starting point...
func main() {
	databases.Migrate(&models.User{}, &models.Comment{}, &models.Article{}) // creating tables

	router := httprouter.New()

	routes.GetRoutes(router)
	routes.GetAuthRoutes(router)
	routes.GetUserRoutes(router)
	routes.GetAdminRoutes(router)

	router.ServeFiles("/static/*filepath", http.Dir("public/"))
	router.ServeFiles("/libs/*filepath", http.Dir("node_modules/"))
	panic(http.ListenAndServe(":8080", router))

}
