// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package
package controllers

import (
	"axeman/libs/layout"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Struct type authController - authController type
type authController struct {
	Layout *layout.LayoutController
}

// AuthController function - returns an initialized pointer of authController
func AuthController() *authController {
	return &authController{}
}

// Auth method -
func (ac *authController) Login(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	ac.Layout.Renderer(res,
		"layout",
		struct{ PageTitle string }{"Login"},
		"views/Layout.html", "views/header.html", "views/index/index.html")
}

// About method -
func (ac *authController) Logout(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	ac.Layout.Renderer(res,
		"layout",
		struct{ PageTitle string }{"Logout"},
		"views/Layout.html", "views/header.html", "views/index/index.html")
}

// Mission method -
func (ac *authController) Signup(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	ac.Layout.Renderer(res,
		"layout",
		struct{ PageTitle string }{"Signup"},
		"views/Layout.html", "views/header.html", "views/auth/signup.html")
}
