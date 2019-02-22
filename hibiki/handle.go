package hibiki

import (
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
	"github.com/sirupsen/logrus"
)

func handleEvents(handler *httphandler.WebhookHandler) httphandler.EventsHandlerFunc {
	client, err := handler.NewClient()
	if err != nil {
		logrus.Fatal(err)
	}

	return func(events []*linebot.Event, r *http.Request) {
		for _, event := range events {
			go handleWebhookEvent(client, event)
		}
	}
}

func handleWebhookEvent(client *linebot.Client, event *linebot.Event) {
	if event.Type == linebot.EventTypeMessage {
		handleMessageEvent(client, event)
	}
}

func handleMessageEvent(client *linebot.Client, event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		handleTextMessage(client, event, message)
	}
}

func handleTextMessage(client *linebot.Client, event *linebot.Event, message *linebot.TextMessage) {
	message.Text = strings.TrimSpace(message.Text)

	// check if message start with "!"
	if strings.HasPrefix(message.Text, "!") {
		message.Text = message.Text[1:]
		handleCommand(client, event, message)
	}
}
