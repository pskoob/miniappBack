package bot

import (
	"context"
	"database/sql"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pskoob/miniappBack/model"
	"go.uber.org/zap"
)

type Bot struct {
	bot             *tgbotapi.BotAPI
	userUsecase     model.IUserUsecase
	usercardUsecase model.IUserCardUsecase
	cardUsecase     model.ICardUsecase
}

func New(userUsecase model.IUserUsecase, usercardUsecase model.IUserCardUsecase, cardUsecase model.ICardUsecase, apiKey string) (*Bot, error) {
	bot, err := tgbotapi.NewBotAPI(apiKey)
	if err != nil {
		zap.L().Error("error init tg bot", zap.Error(err))
		return &Bot{}, err
	}
	zap.L().Info("bot was inizializated with api key " + apiKey)
	return &Bot{
		bot:             bot,
		userUsecase:     userUsecase,
		usercardUsecase: usercardUsecase,
		cardUsecase:     cardUsecase,
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
				TgID:     sql.NullInt64{Int64: tgID, Valid: true},
				Username: sql.NullString{String: username, Valid: true},
			}
			userID, err := b.userUsecase.CreateUser(ctx, user)
			if err != nil {
				zap.L().Error("error create user", zap.Error(err))
			}

			energyBoosterCard, err := b.cardUsecase.GetCardByName(ctx, model.EnergyBooster)
			if err != nil {
				zap.L().Error("error get energy booster card", zap.Error(err))
			}

			powerClickCard, err := b.cardUsecase.GetCardByName(ctx, model.PowerClick)
			if err != nil {
				zap.L().Error("error get power click card", zap.Error(err))
			}

			autoClickerCard, err := b.cardUsecase.GetCardByName(ctx, model.AutoClicker)
			if err != nil {
				zap.L().Error("error get auto clicker card", zap.Error(err))
			}

			energyBooster := model.UserCard{
				UserID:   userID,
				CardID:   energyBoosterCard.ID,
				CardName: energyBoosterCard.Name,
			}

			powerClick := model.UserCard{
				UserID:   userID,
				CardID:   powerClickCard.ID,
				CardName: powerClickCard.Name,
			}

			autoClicker := model.UserCard{
				UserID:   userID,
				CardID:   autoClickerCard.ID,
				CardName: autoClickerCard.Name,
			}

			_, err = b.usercardUsecase.CreateUserCardByUserID(ctx, energyBooster)
			if err != nil {
				zap.L().Error("error create energy booster user card", zap.Error(err))
			}

			_, err = b.usercardUsecase.CreateUserCardByUserID(ctx, powerClick)
			if err != nil {
				zap.L().Error("error create power click user card", zap.Error(err))
			}

			_, err = b.usercardUsecase.CreateUserCardByUserID(ctx, autoClicker)
			if err != nil {
				zap.L().Error("error create auto clicker user card", zap.Error(err))
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
