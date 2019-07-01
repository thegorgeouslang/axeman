// Adminor: James Mallon <jamesmallondev@gmail.com>
// routes package -
package routes

import (
	. "axeman/controllers"
	. "axeman/middlewares"
	"github.com/julienschmidt/httprouter"
)

func GetAdminRoutes(router *httprouter.Router) {

	ac := AdminController()
	acl := ACLMiddleware
	auth := AuthMiddleware()

	router.GET("/dashboard", auth.CheckLogged(ac.Index))
	router.PATCH("/dashboard/modify-user", auth.CheckLogged(acl.Check(ac.ModifyUser)))
}
