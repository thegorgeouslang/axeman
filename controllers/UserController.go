// Useror: James Mallon <jamesmallondev@gmail.com>
// controllers package
package controllers

import (
	log "axeman/libs/logger"
	"axeman/models"
	"axeman/models/dao"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Struct type userController - userController type
type userController struct{}

// UserController function - returns an initialized pointer of userController
func UserController() *userController {
	return &userController{}
}

// Mission method -
func (uc *userController) CreateUser(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	req.ParseForm()
	user := models.User{Email: req.Form.Get("email"), Password: req.Form.Get("password")}

	userDAO := dao.UserDAO()
	r := userDAO.InsertUser(&user)
	if r {
		log.It.WriteLog("alert", "The user wasn't created", log.It.GetTraceMsg())
	}
	http.Redirect(res, req, "/dashboard", 301)
}
