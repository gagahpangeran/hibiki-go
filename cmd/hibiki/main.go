package main

import (
	"github.com/ridho9/hibiki-go/command"
	"github.com/ridho9/hibiki-go/hibiki"
)

func main() {
	hibiki.StartApp(command.DefaultRegistry)
}
