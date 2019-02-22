package command

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/ridho9/hibiki-go/hibiki"
	"github.com/sirupsen/logrus"
)

// DefaultRegistry is the global default command registry
var DefaultRegistry = hibiki.NewRegistry()

func init() {
	DefaultRegistry.AddCommand(createCallCommand())
	logrus.Println("Registered 'call' command to default registry.")
}

func createCallCommand() *hibiki.Command {
	return hibiki.NewCommand(
		"call",
		hibiki.NewDefaultRoute(
			"Usage: !call\nA simple call",
			callDefaultRouteHandler,
		),
	)
}

func callDefaultRouteHandler(req *hibiki.Request, res *hibiki.Response) error {
	res.AddMessage(
		linebot.NewTextMessage("Roger, Hibiki, heading out!\n\nI'll never forget Tenshi..."),
	)
	if req.Text != "" {
		res.AddMessage(
			linebot.NewTextMessage(req.Text),
		)
	}
	err := res.ReplyMessage()

	return err
}
