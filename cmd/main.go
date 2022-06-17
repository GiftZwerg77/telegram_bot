package main

import (
	"log"
  "os"
  "fmt"
  tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

)
var TOKEN string
var productServiece *product.Serviece

func main() {
  TOKEN  := os.Getenv("TOKEN")
  fmt.Println(TOKEN)
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
  wh, _ := tgbotapi.NewWebhook("")
  _, err = bot.Request(wh)
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			if update.Message.IsCommand(){
        prepareCommand(bot,update.Message)
      }else{
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			  bot.Send(msg)
      }

		}
	}
}

func prepareCommand(bot *tgbotapi.BotAPI,msg *tgbotapi.Message){
  switch msg.Command(){
    case "help":
      helpCommand(bot, msg)
    case "list":
      listCommand(bot, msg)
    case "start":
      startCommand(bot, msg)
    default:
      bot.Send(tgbotapi.NewMessage(msg.Chat.ID,msg.Command()+" - this is unknown command"))
  }

}

func helpCommand(bot *tgbotapi.BotAPI,msg *tgbotapi.Message){
  answ := tgbotapi.NewMessage(msg.Chat.ID, "this is trade bot\n")
  bot.Send(answ)
}
func listCommand(bot *tgbotapi.BotAPI,msg *tgbotapi.Message){
  answ := tgbotapi.NewMessage(msg.Chat.ID, "show list\n")
  bot.Send(answ)
}
func startCommand(bot *tgbotapi.BotAPI,msg *tgbotapi.Message){
  answ := tgbotapi.NewMessage(msg.Chat.ID, "START\n")
  bot.Send(answ)
}
