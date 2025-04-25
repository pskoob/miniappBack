package handler

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pskoob/miniappBack/internal/api/definition"
	"github.com/pskoob/miniappBack/internal/api/server/restapi/api"
	"github.com/pskoob/miniappBack/model"
	"go.uber.org/zap"
)

func (h *Handler) SaveUserProgressHandler(req api.SaveProgressParams) middleware.Responder {
	zap.L().Info("save progress request")

	ctx := req.HTTPRequest.Context()

	fmt.Println("1")
	user, err := h.userUsecase.GetUserByTgID(ctx, *req.Progress.TgID)
	if err != nil {
		zap.L().Error("no such user", zap.Error(err))
		return api.NewSaveProgressBadRequest().WithPayload(&definition.Error{
			Message: &model.NoSuchUser,
		})
	}

	fmt.Println("2")

	userCards, err := h.userCardUsecase.GetUserCardsByUserID(ctx, user.ID)
	if err != nil {
		zap.L().Error("no such user", zap.Error(err))
		return api.NewSaveProgressBadRequest().WithPayload(&definition.Error{
			Message: &model.NoSuchUser,
		})
	}

	fmt.Println("3")

	for i := range userCards {
		if userCards[i].CardName == model.AutoClicker && !userCards[i].AutoClicker.Bool {
			autoClickerUserCard := model.UserCard{
				UserID:      userCards[i].UserID,
				CardID:      userCards[i].CardID,
				CardName:    userCards[i].CardName,
				AutoClicker: sql.NullBool{Bool: true, Valid: true},
			}
			err := h.userCardUsecase.UpdateUserCardByUserID(ctx, autoClickerUserCard)
			if err != nil {
				zap.L().Error("error update autor clicker user card", zap.Error(err))
				return api.NewSaveProgressInternalServerError()
			}
		}
		fmt.Println("4")

		if userCards[i].CardName == model.EnergyBooster {
			energyBoosterUserCard := model.UserCard{
				UserID:       userCards[i].UserID,
				CardID:       userCards[i].CardID,
				CardName:     userCards[i].CardName,
				CurrentLevel: *req.Progress.UpgradeEnergy,
				UpdatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
			}
			err := h.userCardUsecase.UpdateUserCardByUserID(ctx, energyBoosterUserCard)
			if err != nil {
				zap.L().Error("error update energy booster user card", zap.Error(err))
				return api.NewSaveProgressInternalServerError()
			}
		}
		fmt.Println("5")

		if userCards[i].CardName == model.PowerClick {
			powerClickUserCard := model.UserCard{
				UserID:       userCards[i].UserID,
				CardID:       userCards[i].CardID,
				CardName:     userCards[i].CardName,
				CurrentLevel: *req.Progress.UpgradeLevel,
				UpdatedAt:    sql.NullTime{Time: time.Now(), Valid: true},
			}
			err := h.userCardUsecase.UpdateUserCardByUserID(ctx, powerClickUserCard)
			if err != nil {
				zap.L().Error("error update power click user card", zap.Error(err))
				return api.NewSaveProgressInternalServerError()
			}
		}
		fmt.Println("6")
	}

	fmt.Println("7")

	err = h.userUsecase.UpdateUserBalance(ctx, user.TgID.Int64, *req.Progress.ClickCount)
	if err != nil {
		zap.L().Error("error update user balance", zap.Error(err))
		return api.NewSaveProgressInternalServerError()
	}
	return api.NewSaveProgressOK().WithPayload(&definition.Error{
		Message: &model.UserDataSaved,
	})
}

func (h *Handler) GetUserProgressHandler(req api.GetUserProgressParams) middleware.Responder {
	zap.L().Info("get user progress request")

	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("error get user by tg id", zap.Error(err))
		return api.NewGetUserProgressInternalServerError()
	}

	return api.NewGetUserProgressOK().WithPayload(&definition.User{
		TgID:       &user.TgID.Int64,
		ClickCount: &user.Balance,
	})
}
