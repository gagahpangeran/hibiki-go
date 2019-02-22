package hibiki

import (
	"fmt"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/sirupsen/logrus"
)

func handleCommand(client *linebot.Client, event *linebot.Event, message *linebot.TextMessage) {
	if len(message.Text) == 0 {
		return
	}

	_, route, rest := globalRegistry.MatchPrefix(message.Text)

	if route == nil {
		commandSplitted := strings.SplitN(message.Text, " ", 2)
		errMessage := fmt.Sprintf("Can't find command %#v", commandSplitted[0])
		logrus.Error(errMessage)
		client.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage(errMessage),
		).Do()
		return
	}

	request := newRequest(event, rest)
	response := newResponse(client, event)

	err := route.handle(request, response)
	if err != nil {
		logrus.Error(message.Text, err)
		client.ReplyMessage(
			event.ReplyToken,
			linebot.NewTextMessage("Ugh..."),
		).Do()
	}
}
