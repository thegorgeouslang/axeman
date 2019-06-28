// Author: James Mallon <jamesmallondev@gmail.com>
// routes package -
package routes

import (
	"axeman/controllers"
	"axeman/middlewares"
	"github.com/julienschmidt/httprouter"
)

func GetAuthRoutes(router *httprouter.Router) {

	ac := controllers.AuthController()
	auth := middlewares.AuthMiddleware()

	router.GET("/login", auth.CheckNonLogged(ac.Login))
	router.GET("/logout", auth.CheckLogged(ac.Logout))
	router.GET("/signup", auth.CheckNonLogged(ac.Signup))
}
