package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pskoob/miniappBack/internal/api/definition"
	"github.com/pskoob/miniappBack/internal/api/server/restapi/api"
	"github.com/pskoob/miniappBack/model"
	"go.uber.org/zap"
)

func (h *Handler) SaveUserProgressHandler(req api.SaveProgressParams) middleware.Responder {
	zap.L().Info("save progress request")

	ctx := req.HTTPRequest.Context()

	err := h.userUsecase.SaveProgressByTgID(ctx, *req.Progress.HasAutoClicker, *req.Progress.UpgradeLevel, *req.Progress.UpgradeEnergy, *req.Progress.ClickCount, *req.Progress.TgID)
	if err != nil {
		zap.L().Error("error save user data", zap.Error(err))
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
		TgID:           &user.TgID.Int64,
		ClickCount:     &user.Balance.Int64,
		HasAutoClicker: &user.AutoClicker,
		UpgradeLevel:   &user.ClickBooster.Int64,
		UpgradeEnergy:  &user.EnergyBooster.Int64,
	})
}
