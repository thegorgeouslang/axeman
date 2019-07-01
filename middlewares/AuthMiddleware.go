// Author: James Mallon <jamesmallondev@gmail.com>
// middlewares package -
package middlewares

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Struct type authMiddleware -
type authMiddleware struct{}

// AuthMiddleware function - check user authentication
func AuthMiddleware() *authMiddleware {
	return &authMiddleware{}
}

// Check method - check user authorization
func (this *authMiddleware) CheckNonLogged(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("Checking if guest")
		next(w, r, p)
	}
}

// Check method - check user authorization
func (this *authMiddleware) CheckLogged(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("Checking if logged user")
		next(w, r, p)
	}
}
