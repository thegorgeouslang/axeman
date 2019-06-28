// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package
package controllers

import (
	"axeman/controllers/helpers"
	"axeman/libs/layout"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Struct type indexController - indexController type
type indexController struct {
	Layout *layout.LayoutController
}

// IndexController function - returns an initialized pointer of indexController
func IndexController() *indexController {
	return &indexController{}
}

// Index method -
func (ic *indexController) Index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	helpers.FlashMessenger().Set(&w, "Hey bro, what's up?!")
	ic.Layout.Renderer(w,
		"layout",
		struct{ PageTitle string }{"Index"},
		"views/layout.html", "views/header.html", "views/index/index.html")
}

// About method -
func (ic *indexController) About(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fm := helpers.FlashMessenger().Get(&w, r)
	ic.Layout.Renderer(w,
		"layout",
		struct {
			PageTitle    string
			FlashMessage string
		}{"About", fm},
		"views/layout.html", "views/header.html", "views/index/index.html")
}

// Mission method -
func (ic *indexController) Mission(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ic.Layout.Renderer(w,
		"layout",
		struct{ PageTitle string }{"Mission"},
		"views/layout.html", "views/header.html", "views/index/index.html")
}

// ContactUs method -
func (ic *indexController) ContactUs(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ic.Layout.Renderer(w,
		"layout",
		struct{ PageTitle string }{"Contact Us"},
		"views/layout.html", "views/header.html", "views/index/index.html")
}
