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

// CreatUser method -
func (this *userController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	user := models.User{Email: r.Form.Get("email"), Password: r.Form.Get("password")}

	userDAO := dao.UserDAO()
	e := userDAO.InsertUser(&user)

	if e != nil {
		log.It.WriteLog("error", e.Error(), log.It.GetTraceMsg())
	}

	http.Redirect(w, r, "/dashboard", 302)
}
