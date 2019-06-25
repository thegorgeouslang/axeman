// Useror: James Mallon <jamesmallondev@gmail.com>
// routes package -
package routes

import (
	"axeman/controllers"
	"github.com/julienschmidt/httprouter"
)

func GetUserRoutes(router *httprouter.Router) {

	router.POST("/user/create", controllers.UserController().CreateUser)
}
