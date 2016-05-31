package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(telegramkey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		switch {
		case update.Message != nil:
			/*
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "hello")
				bot.Send(msg)
			*/
			break
		case update.InlineQuery != nil:
			if len([]rune(update.InlineQuery.Query)) >= 3 {
				res, err := http.Get("http://openstreetmap.ru/api/search?q=" + update.InlineQuery.Query + "&st=&accuracy=1&cnt=5&stype=all&lat=" + strconv.FormatFloat(update.InlineQuery.Location.Latitude, 'f', 7, 64) + "&lon=" + strconv.FormatFloat(update.InlineQuery.Location.Longitude, 'f', 7, 64))

				if err != nil {
					panic(err)
				}

				body, err := ioutil.ReadAll(res.Body)

				resp := new(Places)
				var resources []interface{}
				err = json.Unmarshal(body, &resp)

				if err != nil {
					log.Printf("whoops:", err)
				}

				if len(resp.Matches) != 0 {
					for k, i := range resp.Matches {
						log.Printf(i.Name, i.FullName)

						resloc := InlineQueryResultVenue{}
						resloc.ID = strconv.Itoa(k)
						if i.Name == "" {
							resloc.Title = "Имя не задано"
						} else {
							resloc.Title = i.Name
						}
						resloc.Type = "venue"
						resloc.Longitude = i.Lon
						resloc.Latitude = i.Lat
						resloc.Address = i.FullName

						xt, yt := tilenumber(i.Lat, i.Lon, 19)
						resloc.ThumbURL = "https://c.tile.openstreetmap.org/19/" + xt + "/" + yt + ".png"
						resloc.InputMessageContent = tgbotapi.InputVenueMessageContent{Latitude: i.Lat, Longitude: i.Lon, Title: i.Name, Address: i.FullName}

						resources = append(resources, resloc)
					}
				} else {
					resloc := tgbotapi.InlineQueryResultArticle{}
					resloc.ID = update.InlineQuery.ID
					resloc.Type = "article"
					resloc.Title = "Ничего не найдено"
					resloc.Description = "По вашему запросу '" + update.InlineQuery.Query + "' ничего не найдено. Измените запрос"
					resloc.InputMessageContent = tgbotapi.InputTextMessageContent{Text: "По вашему запросу ничего не найдено"}
					resources = append(resources, resloc)
				}
				ic := tgbotapi.InlineConfig{}
				ic.InlineQueryID = update.InlineQuery.ID
				ic.Results = resources
				bot.AnswerInlineQuery(ic)
			}
			break
		}
	}
}
