package tgbot

import (
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// @server_telega_bot
const (
	TOKEN  = "2043277522:AAH01ln0ntnz3aat6z0QgHS5F0FHC0Ypf5Y"
	chatId = "-1001783728948"
)

var (
	bot *tgbotapi.BotAPI
	u   tgbotapi.UpdateConfig
)

func StartBot() error {
	bot, _ = tgbotapi.NewBotAPI(TOKEN)

	u = tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return nil
}

func SendMessage(msg string) /* (*http.Response, error)*/ {
	bot.Send(tgbotapi.NewMessageToChannel(chatId, msg))
}

//Gin
func Router() {
	router := gin.Default()
	router.GET("/send", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sent",
		})
	})
	router.Run()
}
