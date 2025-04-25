package usecase

import (
	"context"

	"github.com/pskoob/miniappBack/model"
)

type Usecase struct {
	costCardRepository model.ICostCardRepository
}

func New(costCardRepository model.ICostCardRepository) model.ICostCardUsecase {
	return &Usecase{
		costCardRepository: costCardRepository,
	}
}

func (u *Usecase) GetCostCardByCardID(ctx context.Context, cardID int64) (model.CostCard, error) {
	return u.costCardRepository.GetCostCardByCardID(ctx, cardID)
}
