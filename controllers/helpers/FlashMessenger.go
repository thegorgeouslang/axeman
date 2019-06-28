// Author: James Mallon <jamesmallondev@gmail.com>
// Main package
package helpers

import (
	log "axeman/libs/logger"
	"encoding/base64"
	"net/http"
	"time"
)

// Struct type FlashMessenger -
type flashMessenger struct{}

// FlashMessenger function - returns an initialized object of flashMessenger
func FlashMessenger() *flashMessenger {
	return &flashMessenger{}
}

//  Set function - it creates the cookie and set the flash message
func (fm *flashMessenger) Set(w *http.ResponseWriter, message string) {
	c := http.Cookie{ // creates the cookie
		Name:  "flashmessenger",
		Value: base64.URLEncoding.EncodeToString([]byte(message)),
	}
	http.SetCookie(*w, &c)
}

// Get method - it prints the message stored in the cookie
func (fm *flashMessenger) Get(w *http.ResponseWriter, r *http.Request) string {
	c, err := r.Cookie("flashmessenger")
	if err != nil { // check for errors - check if the cookie with the specific name is setted
		log.It.WriteLog("error", err.Error(), log.It.GetTraceMsg())
		return ""
	} else {
		rc := http.Cookie{
			Name:    "flashmessenger",
			MaxAge:  -1,              // a negative number means it's already expired
			Expires: time.Unix(1, 0), // an expiration value that's in the past - already expired
		}
		val, err := base64.URLEncoding.DecodeString(c.Value) // decode the value of the cookie
		http.SetCookie(*w, &rc)                              // override the values of the cookie setted in the setMessage function
		if err != nil {
			log.It.WriteLog("error", err.Error(), log.It.GetTraceMsg())
			return ""
		}
		return string(val)
	}
}
