// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package
package controllers

import (
	l "axeman/libs/layout"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Struct type authController - authController type
type authController struct {
	layout l.LayoutHelper
}

// AuthController function - returns an initialized pointer of authController
func AuthController() *authController {
	return &authController{}
}

// Auth method -
func (this *authController) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	this.layout.Render(w,
		"layout",
		struct{ PageTitle string }{"Login"},
		"templates/layout.gohtml", "templates/header.gohtml", "templates/index/index.gohtml")
}

// About method -
func (this *authController) Logout(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	this.layout.Render(w,
		"layout",
		struct{ PageTitle string }{"Logout"},
		"templates/layout.gohtml", "templates/header.gohtml", "templates/index/index.gohtml")
}

// Mission method -
func (this *authController) Signup(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	this.layout.Render(w,
		"layout",
		struct{ PageTitle string }{"Signup"},
		"templates/layout.gohtml", "templates/header.gohtml", "templates/auth/signup.gohtml")
}
