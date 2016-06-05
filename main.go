package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"gopkg.in/telegram-bot-api.v4"
)

const version = "1.1.3"

func main() {
	config := new(IniConf)
	config.CheckAndLoadConf("config" + string(os.PathSeparator) + "opsconfig.ini")
	telegramkey := config.GetStringKey("", "telegramkey")

	bot, err := tgbotapi.NewBotAPI(telegramkey)
	if err != nil {
		log.Panic("Wrong key:", telegramkey, err)
	}

	bot.Debug = config.GetBoolKey("", "debug")

	// Авторизация бота
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	// Бесконечно ждем апдейтов от сервера
	for update := range updates {
		switch {
		// Пришло обычное сообщение
		case update.Message != nil:
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi!\n\nI'm inline OSM POI search bot v."+version+", type @"+bot.Self.UserName+" in message field")
			bot.Send(msg)
			break
		// Пришел inline запрос
		case update.InlineQuery != nil:
			// Длина запроса >=3?
			if len([]rune(update.InlineQuery.Query)) >= 3 {
				// Обращаемся к osm.ru
				res, err := http.Get("http://openstreetmap.ru/api/search?q=" + update.InlineQuery.Query + "&st=&accuracy=1&cnt=5&stype=all&lat=" + strconv.FormatFloat(update.InlineQuery.Location.Latitude, 'f', 7, 64) + "&lon=" + strconv.FormatFloat(update.InlineQuery.Location.Longitude, 'f', 7, 64))

				if err != nil {
					panic(err)
				}

				body, err := ioutil.ReadAll(res.Body)

				resp := new(Places)

				// Разбираем ответ
				err = json.Unmarshal(body, &resp)

				if err != nil {
					log.Printf("whoops:", err)
				}

				var resources []interface{}

				// Если ответ не нулевой
				if (len(resp.Matches) != 0) && (resp.Search == update.InlineQuery.Query) {
					for k, i := range resp.Matches {
						// Формируем меню venue с полученными пои
						title := "Имя не задано"
						if i.Name != "" {
							title = i.Name
						}
						xt, yt := tilenumber(i.Lat, i.Lon, 19)

						resources = append(resources,
							InlineQueryResultVenue{
								Type:      "venue",
								ID:        strconv.Itoa(k),
								Latitude:  i.Lat,
								Longitude: i.Lon,
								Title:     title + " : " + strconv.FormatInt(Round(Distance(i.Lat, i.Lon, update.InlineQuery.Location.Latitude, update.InlineQuery.Location.Longitude)), 10) + "м",
								Address:   i.FullName,
								InputMessageContent: tgbotapi.InputVenueMessageContent{
									Latitude:  i.Lat,
									Longitude: i.Lon,
									Title:     i.Name,
									Address:   i.FullName},
								ThumbURL: "https://c.tile.openstreetmap.org/19/" + xt + "/" + yt + ".png"})
					}
				} else {
					// Формируем плашку "ничего не найдено"
					resources = append(resources,
						tgbotapi.InlineQueryResultArticle{
							Type:  "article",
							ID:    update.InlineQuery.ID,
							Title: "Ничего не найдено",
							InputMessageContent: tgbotapi.InputTextMessageContent{
								Text: "По вашему запросу ничего не найдено"},
							Description: "По вашему запросу '" + update.InlineQuery.Query + "' ничего не найдено. Измените запрос"})
				}
				// Отправляем меню пользователю
				bot.AnswerInlineQuery(
					tgbotapi.InlineConfig{
						InlineQueryID: update.InlineQuery.ID,
						Results:       resources})
			}
			break
		}
	}
}
