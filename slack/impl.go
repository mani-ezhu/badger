package slack

import (
	"badger/config"
	"badger/constants"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// SlackAlert ...
type SlackAlert struct {
	URL      string
	Username string
	Appname  string
}

// Message ...
type Message struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

// InitSlackAlert ...
// function will init required config on struct
func (s *SlackAlert) InitSlackAlert() {
	// updating on struct
	s.URL = config.Config.HookUrl
	s.Username = config.Config.UserName
	s.Appname = config.Config.Channel
	return
}

// TriggerAlert ...
// function will trigger slack alert configured channel
func (s *SlackAlert) TriggerAlert(data Message) (err error) {
	for {
		// creating byte from message payload
		var msgByte []byte
		if msgByte, err = json.Marshal(data); err != nil {
			fmt.Println("error occured while marshalling slack message", err)
			break
		}
		// call for init config
		s.InitSlackAlert()
		// creating pay load
		payload := make(map[string]interface{})
		payload[constants.UserName] = s.Username
		payload[constants.Text] = fmt.Sprintf("%v", string(msgByte))
		payload[constants.Channel] = s.Appname
		// marshalling payload
		var byteData []byte
		if byteData, err = json.Marshal(payload); err != nil {
			fmt.Println("error occured while marshalling slack payload", err)
			break
		}
		// creating http request to send notification
		client := new(http.Client)
		var req *http.Request
		if req, err = http.NewRequest(http.MethodPost, s.URL, bytes.NewBuffer(byteData)); err != nil {
			fmt.Println("error while creating http request", err)
			break
		}
		// making http call to slack
		var resp *http.Response
		if resp, err = client.Do(req); err != nil {
			fmt.Println("error while sending alert to slack", err)
			return
		}
		// checking for http status
		if resp.StatusCode != http.StatusOK {
			err = errors.New("error occured while sending slack notification")
			return
		}
		break
	}
	return
}
