package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Config struct {
	BotToken string `json:"botToken"`
	BanTime  int64  `json:"banTime"`
}

func ReadConfig() (result Config) {
	file, openErr := os.Open("config.json")
	if openErr != nil {
		log.Fatalf("Could not open config file: %v", openErr)
	}
	bytes, readErr := ioutil.ReadAll(file)
	if readErr != nil {
		log.Fatalf("Could not read config file: %v", readErr)
	}
	unmarshErr := json.Unmarshal(bytes, &result)
	if unmarshErr != nil {
		log.Fatalf("Could not unmarshal JSON config file: %v", unmarshErr)
	}
	return
}
func main() {
	config := ReadConfig()
	bot, err := tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		go func() {
			if update.Message != nil {
				if strings.ContainsAny(update.Message.Text, "安吧八爸百北不大岛的弟地东都对多儿二方港哥个关贵国过海好很会家见叫姐京九可老李零六吗妈么没美妹们名明哪那南你您朋七起千去人认日三上谁什生师十识是四他她台天湾万王我五西息系先香想小谢姓休学也一亿英友月Z再张这中字") {
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
				}
			}

		}()
	}
}
