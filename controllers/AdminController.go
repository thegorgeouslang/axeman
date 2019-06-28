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
	Layout *layout.LayoutController
}

// AdminController function - returns an initialized pointer of adminController
func AdminController() *adminController {
	return &adminController{}
}

// Admin method -
func (ac *adminController) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//helpers.FlashMessenger().Get(w, r)
	ac.Layout.Renderer(w,
		"layout",
		struct{ PageTitle string }{"Admin"},
		"views/admin/dashboard.html", "views/admin/header.html", "views/admin/index.html")
}

// ModifyUser method - implements users modifications by authorized users
func (ac *adminController) ModifyUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ac.Layout.Renderer(w,
		"layout",
		struct{ PageTitle string }{"Admin"},
		"views/admin/dashboard.html", "views/admin/header.html", "views/admin/index.html")
}
