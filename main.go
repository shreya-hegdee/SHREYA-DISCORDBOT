package main

import (
	"fmt"

	"github.com/shreya-hegdee/shreya-discordbot/bot"
	"github.com/shreya-hegdee/shreya-discordbot/config"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	bot.Start()

	<-make(chan struct{})
	return
}
