package telegram

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func botApi() *tgbotapi.BotAPI {
	botApi, err := tgbotapi.NewBotAPI(global.GVA_CONFIG.Telegram.BotToken)
	if err != nil {
		fmt.Println(err.Error())
	}
	return botApi

}

func SendTgMsg(msg string) {
	text := tgbotapi.NewMessage(global.GVA_CONFIG.Telegram.ChannelId, msg)
	botApi().Send(text)
}
