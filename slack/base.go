package slack

// SlackInterface ...
type SlackInterface interface {
	InitSlackAlert()
	TriggerAlert(data Message) error
}

// GetSlack ...
// function for get new slack alert
func GetSlack() SlackInterface {
	return new(SlackAlert)
}
