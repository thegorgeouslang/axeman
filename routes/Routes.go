// Author: James Mallon <jamesmallondev@gmail.com>
// routers package - this package deal with routes
package routes

import (
	"axeman/controllers"
	"github.com/julienschmidt/httprouter"
)

// GetRoutes function -
func GetRoutes(router *httprouter.Router) {

	router.GET("/", controllers.IndexController().Index)
	router.GET("/home", controllers.IndexController().Index)
	router.GET("/about", controllers.IndexController().About)
	router.GET("/mission", controllers.IndexController().Mission)
	router.GET("/contactus", controllers.IndexController().ContactUs)
}
