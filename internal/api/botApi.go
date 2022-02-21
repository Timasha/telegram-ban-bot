package api

import (
	"telegram-ban-bot/internal/utils"
	"unicode"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update, config utils.Config) {
	//if strings.ContainsAny(update.Message.Text, "安吧八爸百北不大岛的弟地东都对多儿二方港哥个关贵国过海好很会家见叫姐京九可老李零六吗妈么没美妹们名明哪那南你您朋七起千去人认日三上谁什生师十识是四他她台天湾万王我五西息系先香想小谢姓休学也一亿英友月Z再张这中字") {
	for _, symbol := range update.Message.Text {
		if unicode.Is(unicode.Han, symbol) {
			deleteMsgConf := tgbotapi.DeleteMessageConfig{
				ChannelUsername: update.Message.Chat.UserName,
				ChatID:          update.Message.Chat.ID,
				MessageID:       update.Message.MessageID,
			}
			banMemberConf := tgbotapi.BanChatMemberConfig{
				ChatMemberConfig: tgbotapi.ChatMemberConfig{
					ChatID:             update.Message.Chat.ID,
					SuperGroupUsername: "",
					ChannelUsername:    update.Message.Chat.UserName,
					UserID:             update.Message.From.ID,
				},
				UntilDate:      config.BanTime,
				RevokeMessages: true,
			}
			bot.Send(deleteMsgConf)
			bot.Send(banMemberConf)
			return
		}
	}
}
