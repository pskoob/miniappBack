package usecase

import (
	"context"

	"github.com/pskoob/miniappBack/model"
)

type Usecase struct {
	cardRepository model.ICardRepository
}

func New(cardRepository model.ICardRepository) model.ICardUsecase {
	return &Usecase{
		cardRepository: cardRepository,
	}
}

func (u *Usecase) GetCardByID(ctx context.Context, cardID int64) (model.Card, error) {
	return u.cardRepository.GetCardByID(ctx, cardID)
}

func (u *Usecase) GetCardByName(ctx context.Context, cardName string) (model.Card, error) {
	return u.cardRepository.GetCardByName(ctx, cardName)
}
