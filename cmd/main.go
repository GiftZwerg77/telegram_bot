package main

import (
	"log"
  "os"
  tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
  "github.com/GiftZwerg77/telegram_bot/internal/serviece/product"
  "github.com/GiftZwerg77/telegram_bot/internal/app/commands"

)
var TOKEN string


func main() {

  TOKEN  := os.Getenv("TOKEN")

  bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}

  productServiece := product.NewServiece()

	bot.Debug = true
  wh, _ := tgbotapi.NewWebhook("")
  _, err = bot.Request(wh)
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

  commander := commands.NewCommander(bot,productServiece)

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			if update.Message.IsCommand(){
        commander.Prepare(bot,update.Message)
      }else{
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			  bot.Send(msg)
      }

		}
	}
}
