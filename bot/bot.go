package bot

import (
	"github.com/pskoob/miniappBack/model"
)

type Bot struct {
	bot tgbotapi
}

func New(userUsecase model.IUserUsecase, apiKey string)
