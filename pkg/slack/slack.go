package slack

import (
	"fmt"

	slackClient "github.com/ashwanthkumar/slack-go-webhook"
	"github.com/labstack/gommon/log"
)

func Send(webhook string, channel string, message string) {
	log.Debugf("WebHook: %s, Channel: %s, Message: %s", webhook, channel, message)

	//attachment1 := slack.Attachment{}
	//attachment1.AddField(slack.Field{Title: "Author", Value: "Ashwanth Kumar"}).AddField(slack.Field{Title: "Status", Value: "Completed"})
	//attachment1.AddAction(slack.Action{Type: "button", Text: "Book flights 🛫", Url: "https://flights.example.com/book/r123456", Style: "primary"})
	//attachment1.AddAction(slack.Action{Type: "button", Text: "Cancel", Url: "https://flights.example.com/abandon/r123456", Style: "danger"})
	payload := slackClient.Payload{
		Text:     message,
		Username: "dokkery",
		Channel:  channel,
	}
	err := slackClient.Send(webhook, "", payload)
	if len(err) > 0 {
		fmt.Printf("error: %s\n", err)
	}
}
