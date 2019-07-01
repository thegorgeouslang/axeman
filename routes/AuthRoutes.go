// Author: James Mallon <jamesmallondev@gmail.com>
// routes package -
package routes

import (
	. "axeman/controllers"
	. "axeman/middlewares"
	"github.com/julienschmidt/httprouter"
)

func GetAuthRoutes(router *httprouter.Router) {

	ac := AuthController()
	auth := AuthMiddleware()

	router.POST("/login", auth.CheckNonLogged(ac.Login))
	router.POST("/logout", auth.CheckLogged(ac.Logout))
	router.POST("/signup", auth.CheckNonLogged(ac.Signup))
}
