// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package
package controllers

import (
	"axeman/libs/layout"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Struct type authController - authController type
type authController struct{}

// AuthController function - returns an initialized pointer of authController
func AuthController() *authController {
	return &authController{}
}

// Auth method -
func (ic *authController) Login(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	layout.Renderer(res,
		"layout",
		struct{ PageTitle string }{"Login"},
		"views/layout.html", "views/header.html", "views/partials/index.html")
}

// About method -
func (ic *authController) Logout(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	layout.Renderer(res,
		"layout",
		struct{ PageTitle string }{"Logout"},
		"views/layout.html", "views/header.html", "views/partials/index.html")
}

// Mission method -
func (ic *authController) Signup(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	layout.Renderer(res,
		"layout",
		struct{ PageTitle string }{"Signup"},
		"views/layout.html", "views/header.html", "views/partials/index.html")
}
