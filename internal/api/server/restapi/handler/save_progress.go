package handler

import (
	"fmt"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pskoob/miniappBack/internal/api/definition"
	"github.com/pskoob/miniappBack/internal/api/server/restapi/api"
	"github.com/pskoob/miniappBack/model"
	"github.com/pskoob/miniappBack/pkg/jsonwebtoken"
	"github.com/pskoob/miniappBack/pkg/useful"
	"go.uber.org/zap"
)

func (h *Handler) SaveUserProgressHandler(req api.SaveProgressParams) middleware.Responder {
	zap.L().Info("save progress request")

	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("no such user", zap.Error(err))
		return api.NewSaveProgressBadRequest().WithPayload(&definition.Error{
			Message: &model.NoSuchUser,
		})
	}

	clickCount, updatedAt, err := jsonwebtoken.ParseToken(*req.TapTokenBody.TapToken, h.tokenSecretKey)
	if err != nil {
		zap.L().Error(fmt.Sprintf("error parse token, userID: %d", user.ID), zap.Error(err))
		return api.NewSaveProgressBadRequest()
	}

	if user.UpdatedAt.Time.Unix() >= updatedAt {
		zap.L().Error("token was expired", zap.Error(err))
		return api.NewSaveProgressBadRequest()
	}

	balance := clickCount + user.Balance

	err = h.userUsecase.UpdateUserBalance(ctx, user.TgID.Int64, balance)
	if err != nil {
		zap.L().Error("error update user balance", zap.Error(err))
		return api.NewSaveProgressInternalServerError()
	}

	return api.NewSaveProgressOK().WithPayload(&definition.Error{
		Message: useful.StrPtr(strconv.Itoa(int(balance))),
	})

}

func (h *Handler) GetUserHandler(req api.GetUserProgressParams) middleware.Responder {
	zap.L().Info("get user progress request")

	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("error get user by tg id", zap.Error(err))
		return api.NewGetUserProgressInternalServerError()
	}

	return api.NewGetUserProgressOK().WithPayload(&definition.User{
		TgID:     &user.TgID.Int64,
		Balance:  &user.Balance,
		Wallet:   &user.Wallet.String,
		Username: &user.Username.String,
	})
}
