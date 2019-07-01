// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package
package controllers

import (
	l "axeman/libs/layout"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Struct type adminController - adminController type
type adminController struct {
	layout l.LayoutHelper
}

// AdminController function - returns an initialized pointer of adminController
func AdminController() *adminController {
	return &adminController{}
}

// Admin method -
func (this *adminController) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//helpers.FlashMessenger().Get(w, r)
	this.layout.Render(w,
		"layout",
		struct{ PageTitle string }{"Admin"},
		"templates/admin/dashboard.gohtml", "templates/admin/index.gohtml")
}

// ModifyUser method - implements users modifications by authorized users
func (this *adminController) ModifyUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	this.layout.Render(w,
		"layout",
		struct{ PageTitle string }{"Admin"},
		"templates/admin/dashboard.gohtml", "templates/admin/index.gohtml")
}
