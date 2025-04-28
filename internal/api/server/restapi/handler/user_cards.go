package handler

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pskoob/miniappBack/internal/api/definition"
	"github.com/pskoob/miniappBack/internal/api/server/restapi/api"
	"github.com/pskoob/miniappBack/model"
	"go.uber.org/zap"
)

func (h *Handler) UpdateUserCardHandler(req api.UpdateUserCardParams) middleware.Responder {
	zap.L().Info("update user card request")
	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("no such user", zap.Error(err))
		fmt.Println("1", err)
		return api.NewUpdateUserCardBadRequest().WithPayload(&definition.Error{
			Message: &model.NoSuchUser,
		})
	}

	userCards, err := h.userCardUsecase.GetUserCardsByUserID(ctx, user.ID)
	if err != nil {
		zap.L().Error("no such card", zap.Error(err))
		return api.NewUpdateUserCardInternalServerError()
	}

	cardResult := &definition.CardBody{}

	for i := range userCards {
		if req.CardUpdateBody.CardName == model.AutoClicker && userCards[i].AutoClicker.Bool {
			return api.NewUpdateUserCardBadRequest().WithPayload(&definition.Error{
				Message: &model.UserAlreadyHasAutoClicker,
			})
		}

		if req.CardUpdateBody.CardName == userCards[i].CardName {
			costCard, err := h.costCardUsecase.GetCostCardByCardID(ctx, userCards[i].CardID)
			if err != nil {
				zap.L().Error("error get card cost", zap.Error(err))
				return api.NewUpdateUserCardInternalServerError()
			}

			card, err := h.cardUsecase.GetCardByID(ctx, userCards[i].CardID)
			if err != nil {
				zap.L().Error("error get card", zap.Error(err))
				return api.NewUpdateUserCardInternalServerError()
			}

			if userCards[i].CurrentLevel == card.MaxLevel.Int64 {
				return api.NewUpdateUserCardBadRequest().WithPayload(&definition.Error{
					Message: &model.CardHasMaxLvl,
				})
			}

			updatePrice := float64(costCard.BasePrice) * math.Pow(2, float64(userCards[i].CurrentLevel))
			if userCards[i].CurrentLevel == 0 {
				updatePrice = float64(costCard.BasePrice)
			}

			if user.Balance < int64(updatePrice) {
				return api.NewUpdateUserCardBadRequest().WithPayload(&definition.Error{
					Message: &model.BalanceTooLow,
				})
			}

			fmt.Println("int: ", int64(updatePrice), "float: ", updatePrice)

			newUserBalance := user.Balance - int64(updatePrice)

			err = h.userUsecase.UpdateUserBalance(ctx, user.TgID.Int64, newUserBalance)
			if err != nil {
				zap.L().Error("error update user balance", zap.Error(err))
				return api.NewUpdateUserCardInternalServerError()
			}

			if req.CardUpdateBody.CardName == model.AutoClicker {
				userCards[i].AutoClicker.Bool = true
				err := h.userCardUsecase.UpdateUserCardByUserID(ctx, userCards[i])
				if err != nil {
					return api.NewUpdateUserCardInternalServerError()
				}
			}

			if req.CardUpdateBody.CardName == model.PowerClick {
				userCards[i].CurrentLevel = userCards[i].CurrentLevel + 1
				err := h.userCardUsecase.UpdateUserCardByUserID(ctx, userCards[i])
				if err != nil {
					return api.NewUpdateUserCardInternalServerError()
				}
			}

			if req.CardUpdateBody.CardName == model.EnergyBooster {
				userCards[i].CurrentLevel = userCards[i].CurrentLevel + 1
				err := h.userCardUsecase.UpdateUserCardByUserID(ctx, userCards[i])
				if err != nil {
					return api.NewUpdateUserCardInternalServerError()
				}
			}

			cardResult.AutoClicker = userCards[i].AutoClicker.Bool
			cardResult.CardID = userCards[i].CardID
			cardResult.CardName = userCards[i].CardName
			cardResult.CurrentLevel = userCards[i].CurrentLevel
			cardResult.UserID = userCards[i].UserID
			cardResult.UpdatedAt = time.Now().Unix()
			cardResult.CreatedAt = userCards[i].CreatedAt.Unix()

		}

	}

	return api.NewUpdateUserCardOK().WithPayload(cardResult)

}

func (h *Handler) GetUserCardsHandler(req api.GetUserCardsParams) middleware.Responder {
	zap.L().Info("get user cards request")

	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error(fmt.Sprintf("no such user tg id: %d", req.TgID), zap.Error(err))
		return api.NewGetUserCardsBadRequest()
	}

	userCards, err := h.userCardUsecase.GetUserCardsByUserID(ctx, user.ID)
	if err != nil {
		zap.L().Error(fmt.Sprintf("error get user cards, user tg id: %d", req.TgID), zap.Error(err))
		return api.NewGetUserCardsInternalServerError()
	}

	userCardsResult := h.UserCardsToDefinition(ctx, userCards)
	return api.NewGetUserCardsOK().WithPayload(&definition.CardsBody{
		Count:     int64(len(userCardsResult)),
		UserCards: userCardsResult,
	})
}

func (h *Handler) UserCardsToDefinition(ctx context.Context, userCards []model.UserCard) []*definition.CardBody {
	userCardsresult := make([]*definition.CardBody, len(userCards))

	for i := range userCards {
		userCardsresult[i] = &definition.CardBody{
			UserID:       userCards[i].UserID,
			CardID:       userCards[i].CardID,
			CardName:     userCards[i].CardName,
			CurrentLevel: userCards[i].CurrentLevel,
			AutoClicker:  userCards[i].AutoClicker.Bool,
			CreatedAt:    userCards[i].CreatedAt.Unix(),
		}

		if userCards[i].UpdatedAt.Valid {
			userCardsresult[i].UpdatedAt = userCards[i].UpdatedAt.Time.Unix()
		}
	}

	return userCardsresult
}
