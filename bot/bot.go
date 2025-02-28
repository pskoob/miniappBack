package bot

import (
	"context"
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pskoob/miniappBack/model"
	"go.uber.org/zap"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	userUsecase model.IUserUsecase
}

func New(userUsecase model.IUserUsecase, apiKey string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		zap.L().Error("error init tg bot", zap.Error(err))
		return &Bot{}, err
	}
	zap.L().Info("bot was inizializated with api key " + apiKey)
	return &Bot{
		bot:         bot,
		userUsecase: userUsecase,
	}, err
}

func (b *Bot) ListenUpdates() {
	u := tgbotapi.NewUpdate(0)
	u.AllowedUpdates = []string{"messages"}
	u.Timeout = 60
	updates := b.bot.GetUpdatesChan(u)
	zap.L().Info("Start listening for tg messages")

	for update := range updates {
		if err := b.ProcessMessage(update); err != nil {
			zap.L().Error("error process tg messages", zap.Error(err))
		}
	}

}

func (b *Bot) ProcessMessage(update tgbotapi.Update) error {

	ctx := context.Background()

	if update.Message != nil {
		tgID := update.Message.From.ID
		username := update.Message.From.UserName
		_, err := b.userUsecase.GetUserByTgID(ctx, tgID)
		if err != nil {
			user := model.User{
				Name:     sql.NullString{String: username, Valid: true},
				TgID:     sql.NullInt64{Int64: tgID, Valid: true},
				Username: sql.NullString{String: username, Valid: true},
			}
			err = b.userUsecase.CreateUser(ctx, user)
			if err != nil {
				zap.L().Error("error create user", zap.Error(err))
			}
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msgText := "Welcome to the Clicker Game! Click the menu button to start"
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)

				if _, err := b.bot.Send(msg); err != nil {
					zap.L().Error("error send message", zap.Error(err))
				}

			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "I don't know that command")
				if _, err := b.bot.Send(msg); err != nil {
					zap.L().Error("error send message", zap.Error(err))
				}
			}
		}
	}

	return nil
}
