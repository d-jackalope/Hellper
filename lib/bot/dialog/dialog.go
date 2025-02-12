package dialog

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/JackBekket/hellper/lib/bot/command"
	"github.com/JackBekket/hellper/lib/database"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleUpdates(updates <-chan tgbotapi.Update, bot *tgbotapi.BotAPI, comm command.Commander, db_service *database.Service) {
	ai_endpoint := os.Getenv("AI_ENDPOINT") // TODO: should not be here?
	for update := range updates {
		if update.CallbackQuery == nil {

			var group = false
			if update.Message.Chat.ID < 0 {
				group = true
			}
			if group && !strings.Contains(update.Message.Text, bot.Self.UserName) && update.Message.Voice == nil && update.Message.Photo == nil && update.Message.Command() == "" {
				continue
			}

			if group && update.Message.Photo != nil && !strings.Contains(update.Message.Caption, bot.Self.UserName) {
				continue
			} else {
				re := regexp.MustCompile(`@?` + regexp.QuoteMeta(bot.Self.UserName))
				update.Message.Caption = re.ReplaceAllString(update.Message.Caption, "")
				update.Message.Caption = strings.TrimSpace(update.Message.Caption)
			}

			chatID := int64(update.Message.Chat.ID)
			// in memory (cashe) db
			db := comm.GetUsersDb()
			user, ok := db[int64(chatID)]
			if !ok {	// if there are no record in memory
				
				//TODO: Here we should try to fetch user from actual db and put it in cash if found
				ds := db_service
				user_exist_in_db := ds.CheckSession(chatID)

				if user_exist_in_db == true {
					// download user data from database into cashe
					ai_session, _ := ds.GetSession(chatID)
					model := ai_session.Model
					url := ai_session.Endpoint.URL
					user := database.User{
						ID:           chatID,
						Username:     update.Message.From.UserName,
						DialogStatus: 6,
						Admin:        false,
					}
					user.AiSession.GptModel = *model
					user.AiSession.Base_url = url
					api_key, err := ds.GetToken(chatID,1)
					if err != nil {
						log.Println("error getting user api key", err)
					}
					user.AiSession.GptKey = api_key

					//history := ds.GetHistory(chatID)
					//TODO: download and paste messages history to the dialog buffer here

					database.AddUser(user)	// add user to cash db
				}

				// user do not exist nor in cash nor in persistent db
				// then we setup dialog
				comm.AddNewUserToMap(update.Message,ai_endpoint)
				//user.AiSession.Base_url = ai_endpoint 	// TODO: it is hardcode here, need refactor
				//db[chatID] = user
			}
			
			

			if ok {

				if update.Message == nil {
					continue
				}

				if !ok {
					comm.AddNewUserToMap(update.Message,ai_endpoint)	// TODO: is it double? why is it here?
				}
				if ok {

					log.Println("user dialog status:", user.DialogStatus)
					log.Println(user.ID)
					log.Println(user.Username)

					if group && update.Message.Voice != nil && user.DialogStatus != 6 {
						continue
					}
					if update.Message.Text != "" {
						re := regexp.MustCompile(`@?` + regexp.QuoteMeta(bot.Self.UserName))
						update.Message.Text = re.ReplaceAllString(update.Message.Text, "")
					}
					switch user.DialogStatus {

					case 3:
						comm.ChooseModel(update.Message,db_service)
					case 4, 5:
						comm.WrongResponse(update.Message)
					case 6:
						comm.DialogSequence(update.Message, ai_endpoint,db_service)

					}

				}

			} // usual handle end

		} else {
			//here goes the callback logic for inlines
			chatID := int64(update.CallbackQuery.Message.Chat.ID)
			db := comm.GetUsersDb()
			user := db[int64(chatID)]
			ai_endpoint := os.Getenv("AI_ENDPOINT")

			switch user.DialogStatus {
			case 4:
				comm.HandleModelChoose(update.CallbackQuery)
			case 5:
				comm.ConnectingToAiWithLanguage(update.CallbackQuery, ai_endpoint)
				// after successful connection we can save user from cashe to persistent db
				// TODO: find a way to get thread id here
				//db_service.CreateChatSession(update.CallbackQuery.Message.Chat.ID,1,chatID,chatID,user.AiSession.GptModel)
				//req := update.CallbackQuery
				//req_msg :=
				//db_service.UpdateHistory(chatID,1,chatID,chatID,user.AiSession.GptModel,)
				db_service.CreateLSession(chatID,user.AiSession.GptModel)
			}
		}
	} // end of main func
}
