package main

import (
	"fmt"
	"server-bot/tgbot"
)

// @title Server Bot
// @version 1.0
// @description This is a server which sends message to telegram group
// @host localhost:8080
// @BasePath /

func main() {
	m1 := tgbot.NewMessage("low", "allo")
	db := tgbot.MessageDB{}
	db.MDB = append(db.MDB, *m1)
	fmt.Println(db.MDB)

	tgbot.SetMessageDB(&db)

	tgbot.StartBot()
	tgbot.Router()
}
