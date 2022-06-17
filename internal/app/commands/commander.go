package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
import "github.com/GiftZwerg77/telegram_bot/internal/serviece/product"

type Commander struct{
  bot *tgbotapi.BotAPI
  productServiece *product.Serviece
}
func NewCommander(bot *tgbotapi.BotAPI, productServiece *product.Serviece) *Commander{
  return &Commander{ bot : bot,
                     productServiece : productServiece,}
}
func (c *Commander) Prepare(bot *tgbotapi.BotAPI,msg *tgbotapi.Message){
  switch msg.Command(){
    case "help":
      c.help(msg)
    case "list":
      c.list(msg)
    case "start":
      c.start(msg)
    default:
      c.bot.Send(tgbotapi.NewMessage(msg.Chat.ID,msg.Command()+" - this is unknown command"))
  }

}


func (c *Commander) help(msg *tgbotapi.Message){
  answ := tgbotapi.NewMessage(msg.Chat.ID, "this is trade bot\n")
  c.bot.Send(answ)
}
func (c *Commander) list(msg *tgbotapi.Message){
  answ := tgbotapi.NewMessage(msg.Chat.ID, "show list\n")
  c.bot.Send(answ)
}
func (c *Commander) start(msg *tgbotapi.Message){
  answ := tgbotapi.NewMessage(msg.Chat.ID, "START\n")
  c.bot.Send(answ)
}
