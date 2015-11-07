package messenger

// This subsystem is responsible for all realtime user chat. That could be anything
// from connecting to IRC servers or integrating with telegram.org.

import (
	"github.com/twitchyliquid64/CNC/logging"
	"github.com/twitchyliquid64/CNC/config"
  "github.com/Syfaro/telegram-bot-api"
)

// global variables
var gTelegramConnection *tgbotapi.BotAPI
var gTelegramBotUsername string

func Initialise()error{
	logging.Info("messenger", "Initialise()")

  gTelegramBotUsername = config.All().Messenger.TelegramIntegration.BotUsername
  logging.Info("messenger", "Now connecting to telegram.org 'BotFather' using ", gTelegramBotUsername)

  var err error
  gTelegramConnection, err = tgbotapi.NewBotAPI(config.All().Messenger.TelegramIntegration.Token)
  if err != nil {
      logging.Error("messenger", err.Error())
  }else {
    go TelegramMessageHandler()
  }

	return err
}
