package main

import (
	"encoding/json"
	"github.com/baxaa/e-commerce/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
)

type Photo struct {
	Urls Urls `json:"urls"`
}
type Urls struct {
	Regular string `json:"regular"`
}

func GetRandomPhoto() (string, error) {
	url := "https://api.unsplash.com/photos/random?client_id=" + config.Acces_key
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var photo Photo
	err = json.NewDecoder(resp.Body).Decode(&photo)
	if err != nil {
		return "", err
	}
	return photo.Urls.Regular, nil
}

func main() {
	bot, err := tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	count := 0
	done := make(chan bool)

	go func() {
		for update := range updates {
			if update.Message != nil {
				if update.Message.IsCommand() && update.Message.Command() == "image" || update.Message.Text == "image" {
					count++
					photo, err := GetRandomPhoto()
					if err != nil {
						log.Println(err)
					}
					file := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileURL(photo))
					done <- true
					bot.Send(file)
				}
			}
		}
	}()
	go func() {
		for {
			log.Printf("Count: %v", count)
			<-done
		}
	}()
	select {}
}
