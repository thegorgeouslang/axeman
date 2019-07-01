// Author: James Mallon <jamesmallondev@gmail.com>
// routers package - this package deal with routes
package routes

import (
	. "axeman/controllers"
	"github.com/julienschmidt/httprouter"
)

// GetRoutes function -
func GetRoutes(router *httprouter.Router) {

	router.GET("/", IndexController().Index)
	router.GET("/home", IndexController().Index)
	router.GET("/about", IndexController().About)
	router.GET("/mission", IndexController().Mission)
	router.GET("/contactus", IndexController().ContactUs)
}
