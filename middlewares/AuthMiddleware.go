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
func (am *authMiddleware) Check(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("Checking user authentication")
		next(w, r, p)
	}
}
