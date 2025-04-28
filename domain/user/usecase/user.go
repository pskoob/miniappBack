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

func (u *Usecase) CreateUser(ctx context.Context, user model.User) (int64, error) {
	return u.userRepository.CreateUser(ctx, user)
}

func (u *Usecase) GetUserByTgID(ctx context.Context, tgID int64) (model.User, error) {
	return u.userRepository.GetUserByTgID(ctx, tgID)
}

func (u *Usecase) GetUserByReferralLink(ctx context.Context, referralLink string) (model.User, error) {
	return u.userRepository.GetUserByReferralLink(ctx, referralLink)
}

func (u *Usecase) GetUsersByReferralUser(ctx context.Context, userID int64) ([]model.User, error) {
	return u.userRepository.GetUsersByReferralUser(ctx, userID)
}

func (u *Usecase) UpdateUserBalance(ctx context.Context, tgID int64, balance int64) error {
	return u.userRepository.UpdateUserBalance(ctx, tgID, balance)
}

func (u *Usecase) UpdateEnergyByTgID(ctx context.Context, energy int64, tgID int64) error {
	return u.userRepository.UpdateEnergyByTgID(ctx, energy, tgID)
}

func (u *Usecase) UpdateWalletByTgID(ctx context.Context, wallet string, tgID int64) error {
	return u.userRepository.UpdateWalletByTgID(ctx, wallet, tgID)
}

func (u *Usecase) UpdateUserReferralByTgID(ctx context.Context, referralUserID int64, tgID int64) error {
	return u.userRepository.UpdateUserReferralByTgID(ctx, referralUserID, tgID)
}

func (u *Usecase) UpdateTotalCoinEarnedByTgID(ctx context.Context, totalCoinEarned int64, tgID int64) error {
	return u.userRepository.UpdateTotalCoinEarnedByTgID(ctx, totalCoinEarned, tgID)
}

func (u *Usecase) UpdateTotalCoinSentToReferralByTgID(ctx context.Context, totalCoinSentToReferral int64, tgID int64) error {
	return u.userRepository.UpdateTotalCoinSentToReferralByTgID(ctx, totalCoinSentToReferral, tgID)
}
