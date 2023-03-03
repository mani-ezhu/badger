package controller

import (
	"badger/constants"
	"badger/slack"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// SpammerRequest ...
type SpammerRequest struct {
	RecordType    string    `json:"RecordType"`
	Type          string    `json:"Type"`
	TypeCode      int       `json:"TypeCode"`
	Name          string    `json:"Name"`
	Tag           string    `json:"Tag"`
	MessageStream string    `json:"MessageStream"`
	Description   string    `json:"Description"`
	Email         string    `json:"Email"`
	From          string    `json:"From"`
	BouncedAt     time.Time `json:"BouncedAt"`
}

// Spammer ...
// Controller function send spammer notification
func Spammer(r *http.Request) (err error) {
	for {
		// reading request body
		var reqBody []byte
		if reqBody, err = ioutil.ReadAll(r.Body); err != nil {
			fmt.Println("error occured while reading request body", err)
			break
		}
		// check for request body empty or not
		if len(reqBody) == 0 {
			err = errors.New("request body can not be empty")
			break
		}
		// unmarshalling request with request struct
		var spammerReq SpammerRequest
		if err = json.Unmarshal(reqBody, &spammerReq); err != nil {
			fmt.Println("error occured while unmarshalling request object", err)
			break
		}
		// check for given request eligible for send notication
		if spammerReq.Type != constants.SpamNotification {
			err = errors.New("request not eligible for send notification on slack")
			break
		}
		// creating message payload
		var payload slack.Message
		payload.Email = spammerReq.Email
		payload.Message = spammerReq.Description
		// call for send slack alret
		err = slack.GetSlack().TriggerAlert(payload)
		break
	}
	return
}
