// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package
package controllers

import (
	. "axeman/controllers/helpers"
	l "axeman/libs/layout"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Struct type indexController - indexController type
type indexController struct {
	layout l.LayoutHelper
}

// IndexController function - returns an initialized pointer of indexController
func IndexController() *indexController {
	return &indexController{}
}

// Index method -
func (this *indexController) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	FlashMessenger().Set(&w, "Hey bro, what's up?!")
	this.layout.Render(w,
		"layout",
		struct{ PageTitle string }{"Index"},
		"templates/layout.html", "templates/header.html", "templates/index/index.html")
}

// About method -
func (this *indexController) About(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fm := FlashMessenger().Get(&w, r)
	this.layout.Render(w,
		"layout",
		struct {
			PageTitle    string
			FlashMessage string
		}{"About", fm},
		"templates/layout.html", "templates/header.html", "templates/index/index.html")
}

// Mission method -
func (this *indexController) Mission(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	this.layout.Render(w,
		"layout",
		struct{ PageTitle string }{"Mission"},
		"templates/layout.html", "templates/header.html", "templates/index/index.html")
}

// ContactUs method -
func (this *indexController) ContactUs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	this.layout.Render(w,
		"layout",
		struct{ PageTitle string }{"Contact Us"},
		"templates/layout.html", "templates/header.html", "templates/index/index.html")
}
