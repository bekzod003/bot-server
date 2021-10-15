package main

import (
	"fmt"
	"server-bot/tgbot"
)

func main() {
	m1 := tgbot.NewMessage("low", "allo")
	db := tgbot.MessageDB{}
	db.MDB = append(db.MDB, *m1)
	fmt.Println(db.MDB)

	tgbot.SetMessageDB(&db)

	tgbot.StartBot()
	tgbot.Router()
}
