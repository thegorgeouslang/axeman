// Adminor: James Mallon <jamesmallondev@gmail.com>
// routes package -
package routes

import (
	"axeman/controllers"
	"axeman/middlewares"
	"github.com/julienschmidt/httprouter"
)

func GetAdminRoutes(router *httprouter.Router) {

	ac := controllers.AdminController()
	acl := middlewares.ACLMiddleware
	auth := middlewares.AuthMiddleware()

	router.GET("/dashboard", auth.Check(ac.Index))
	router.GET("/dashboard/modify-user", auth.Check(acl.Check(ac.ModifyUser)))
}
