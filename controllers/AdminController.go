// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package
package controllers

import (
	"axeman/libs/layout"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Struct type adminController - adminController type
type adminController struct {
}

// AdminController function - returns an initialized pointer of adminController
func AdminController() *adminController {
	return &adminController{}
}

// Admin method -
func (ac *adminController) Index(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	layout.Renderer(res,
		"layout",
		struct{ PageTitle string }{"Admin"},
		"views/admin/dashboard.html", "views/header.html", "views/admin/index.html")
}
