// Author: James Mallon <jamesmallondev@gmail.com>
// middlewares package - code related to route middlewares
package middlewares

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// Struct type aclMiddleware -
type aclMiddleware struct{}

var ACLMiddleware *aclMiddleware

// init function - data and process initialization
func init() {
	ACLMiddleware = &aclMiddleware{}
}

// Check method - check user authorization
func (this *aclMiddleware) Check(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Println("Checking user authorization...")
		next(w, r, p)
	}
}
