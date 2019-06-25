// Adminor: James Mallon <jamesmallondev@gmail.com>
// routes package -
package routes

import (
	"axeman/controllers"
	"github.com/julienschmidt/httprouter"
)

func GetAdminRoutes(router *httprouter.Router) {

	router.GET("/dashboard", controllers.AdminController().Index)
}
