package hibiki

import "github.com/line/line-bot-sdk-go/linebot"

// Request struct
type Request struct {
	Event *linebot.Event
	Text  string
}

func newRequest(event *linebot.Event, text string) *Request {
	return &Request{
		event,
		text,
	}
}
