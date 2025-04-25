package usecase

import (
	"context"

	"github.com/pskoob/miniappBack/model"
)

type Usecase struct {
	usercardRepository model.IUserCardRepository
}

func New(usercardRepository model.IUserCardRepository) model.IUserCardUsecase {
	return &Usecase{
		usercardRepository: usercardRepository,
	}
}

func (u *Usecase) CreateUserCardByUserID(ctx context.Context, userCard model.UserCard) (int64, error) {
	return u.usercardRepository.CreateUserCardByUserID(ctx, userCard)
}

func (u *Usecase) GetUserCardsByUserID(ctx context.Context, userID int64) ([]model.UserCard, error) {
	return u.usercardRepository.GetUserCardsByUserID(ctx, userID)
}

func (u *Usecase) UpdateUserCardByUserID(ctx context.Context, userCard model.UserCard) error {
	return u.usercardRepository.UpdateUserCardByUserID(ctx, userCard)
}
