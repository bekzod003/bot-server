package tgbot

import (
	"time"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// username = server_telega_bot
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

var msgdb MessageDB

func SetMessageDB(msg *MessageDB) {
	msgdb = *msg
}

// @Summary Send Message
// @Description Send message to telegram group
// @ID send-message
// @Accept json
// @Produce json
// @Router /send [get]

func SendMessage( /*msg string*/ ) /* (*http.Response, error)*/ {
	getMsg := GetMessageByPriority(&msgdb)
	bot.Send(tgbotapi.NewMessageToChannel(chatId, getMsg))
}

//Gin
func Router() {
	router := gin.Default()
	router.GET("/send", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "sent",
		})
		var td time.Duration = time.Second * 5

		time.AfterFunc(td, SendMessage)
	})
	router.Run()
}

func GetMessageByPriority(msg *MessageDB) string {
	for _, val := range msg.MDB {
		if val.Priority == "high" {
			tempTxt := val.Text
			msg.Delete(val)
			return tempTxt
		}
	}
	for _, val := range msg.MDB {
		if val.Priority == "medium" {
			tempTxt := val.Text
			msg.Delete(val)
			return tempTxt
		}
	}
	for _, val := range msg.MDB {
		if val.Priority == "low" {
			tempTxt := val.Text
			msg.Delete(val)
			return tempTxt
		}
	}
	return ""
}
