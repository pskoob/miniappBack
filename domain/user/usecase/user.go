package usecase

import (
	"context"

	"github.com/pskoob/miniappBack/model"
)

type Usecase struct {
	userRepository model.IUserRepository
}

func New(userRepository model.IUserRepository) model.IUserUsecase {
	return &Usecase{
		userRepository: userRepository,
	}
}

func (u *Usecase) CreateUser(ctx context.Context, user model.User) error {
	return u.userRepository.CreateUser(ctx, user)
}

func (u *Usecase) GetUserByTgID(ctx context.Context, tgID int64) (model.User, error) {
	return u.userRepository.GetUserByTgID(ctx, tgID)
}

func (u *Usecase) SaveProgressByTgID(ctx context.Context, autoClicker bool, clickBooster int64, energyBooster int64, balance int64, tgID int64) error {
	return u.userRepository.SaveProgressByTgID(ctx, autoClicker, clickBooster, energyBooster, balance, tgID)
}
