package handler

import (
	"badger/controller"
	"badger/util"
	"net/http"
)

// SpammerHandler ...
// Handler function for send notification based on incoming message
func SpammerHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	for {
		// check for request method
		if r.Method == http.MethodPost {
			// invoking respective controller
			err = controller.Spammer(r)
		}
		break
	}
	util.GenrateResponse(w, err)
	return
}
