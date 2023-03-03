package util

import (
	"badger/constants"
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/cast"
)

// GenrateResponse ...
// function will generate final response
func GenrateResponse(w http.ResponseWriter, err error) {
	w.Header().Set(constants.ConstContentType, constants.ConstApplicationSlashJSON)
	w.WriteHeader(http.StatusOK)
	resp := make(map[string]interface{})
	resp[constants.ConstMessage] = constants.ConstSuccess
	resp[constants.ConstStatusCode] = constants.Const200
	if err != nil {
		w.WriteHeader(constants.Const451)
		resp[constants.ConstMessage] = cast.ToString(err)
		resp[constants.ConstStatusCode] = constants.Const451
	}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
}
