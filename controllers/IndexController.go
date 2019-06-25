// Author: James Mallon <jamesmallondev@gmail.com>
// controllers package
package controllers

import (
	"axeman/libs/layout"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Struct type indexController - indexController type
type indexController struct {
}

// IndexController function - returns an initialized pointer of indexController
func IndexController() *indexController {
	return &indexController{}
}

// Index method -
func (ic *indexController) Index(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	layout.Renderer(res,
		"layout",
		struct{ PageTitle string }{"Index"},
		"views/layout.html", "views/header.html", "views/index/index.html")
}

// About method -
func (ic *indexController) About(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	layout.Renderer(res,
		"layout",
		struct{ PageTitle string }{"About"},
		"views/layout.html", "views/header.html", "views/index/index.html")
}

// Mission method -
func (ic *indexController) Mission(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	layout.Renderer(res,
		"layout",
		struct{ PageTitle string }{"Mission"},
		"views/layout.html", "views/header.html", "views/index/index.html")
}

// ContactUs method -
func (ic *indexController) ContactUs(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	layout.Renderer(res,
		"layout",
		struct{ PageTitle string }{"Contact Us"},
		"views/layout.html", "views/header.html", "views/index/index.html")
}
