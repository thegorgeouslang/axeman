// Author: James Mallon <jamesmallondev@gmail.com>
// routes package -
package routes

import (
	"axeman/controllers"
	"github.com/julienschmidt/httprouter"
)

func GetAuthRoutes(router *httprouter.Router) {

	router.GET("/login", controllers.AuthController().Login)
	router.GET("/logout", controllers.AuthController().Logout)
	router.GET("/signup", controllers.AuthController().Signup)
	router.POST("/signup-process", controllers.AuthController().SignupProcess)
}
