package hibiki

import (
	"fmt"
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot/httphandler"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hibiki running")
	})
}

func StartApp() {
	handler, err := httphandler.New(
		viper.GetString("line.channel_secret"),
		viper.GetString("line.channel_access_token"),
	)
	if err != nil {
		logrus.Fatal("Failed initializing line handler")
	}
	logrus.Println("Initialized line handler")

	handler.HandleEvents(handleEvents(handler))

	http.Handle("/callback", handler)

	addrString := fmt.Sprintf("%v:%v", viper.GetString("hostname"), viper.GetInt("port"))

	logrus.Printf("Running hibiki at %v", addrString)

	if err := http.ListenAndServe(
		addrString,
		nil,
	); err != nil {
		logrus.Fatal(err)
	}
}
