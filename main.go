package main

import (
	"github.com/avestoinn/captchabot/bot"
	"github.com/avestoinn/captchabot/config"
	"github.com/avestoinn/captchabot/messages"
	"github.com/avestoinn/captchabot/models"
)

func main() {
	config.Setup("config.json")
	models.SetupDb()
	messages.InitTranslations()
	bot.Run()
}
