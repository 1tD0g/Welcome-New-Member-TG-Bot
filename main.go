package main

import (
	"fmt"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

const token = ""

func main(){
	bot, err := tg.NewBotAPI(token)
	if err != nil { log.Fatal(err)}
	u, err := bot.GetUpdatesChan(tg.NewUpdate(0))
	if err != nil { log.Fatal(err)}
	for eu := range u {
		if eu.Message.NewChatMembers == nil { continue }
		for _, v := range *eu.Message.NewChatMembers {
			log.Printf("New User @%s in Group ID: %d", v.UserName, eu.Message.Chat.ID)
			msg := tg.NewMessage(eu.Message.Chat.ID, fmt.Sprintf("歡迎新面孔: <a href=\"tg://user?id=%d\">%s %s</a>❗️❗️❗️❗️❗️", v.ID, v.LastName, v.FirstName))
			msg.ParseMode = tg.ModeHTML
			msg.ReplyToMessageID = eu.Message.MessageID
			_, err := bot.Send(msg)
			if err != nil { log.Println("Error Sending Message: ", err)}
		}
	}
}
