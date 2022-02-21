package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
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
