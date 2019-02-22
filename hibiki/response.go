package hibiki

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// Response struct
type Response struct {
	client     *linebot.Client
	replyToken string
	messages   []linebot.SendingMessage
}

// newResponse creates new response object
func newResponse(client *linebot.Client, event *linebot.Event) *Response {
	return &Response{
		client,
		event.ReplyToken,
		make([]linebot.SendingMessage, 0),
	}
}

// AddMessage append message to response message slice
func (r *Response) AddMessage(msg linebot.SendingMessage) {
	r.messages = append(r.messages, msg)
}

func (r *Response) ReplyMessage() error {
	_, err := r.client.ReplyMessage(r.replyToken, r.messages...).Do()
	return err
}
