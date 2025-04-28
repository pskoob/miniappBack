package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

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
		fmt.Println(err)
		return api.NewSaveProgressBadRequest().WithPayload(&definition.Error{
			Message: &model.NoSuchUser,
		})
	}

	clickCount, updatedAt, err := jsonwebtoken.ParseTapToken(*req.TapTokenBody.TapToken, h.tokenSecretKey)
	if err != nil {
		zap.L().Error(fmt.Sprintf("error parse token, userID: %d", user.ID), zap.Error(err))
		return api.NewSaveProgressBadRequest()
	}

	if user.UpdatedAt.Time.Unix() >= updatedAt {
		zap.L().Error("tap token was expired", zap.Error(err))
		return api.NewSaveProgressBadRequest()
	}

	balance := clickCount + user.Balance

	err = h.userUsecase.UpdateUserBalance(ctx, user.TgID.Int64, balance)
	if err != nil {
		zap.L().Error("error update user balance", zap.Error(err))
		return api.NewSaveProgressInternalServerError()
	}

	err = h.userUsecase.UpdateTotalCoinEarnedByTgID(ctx, balance, user.TgID.Int64)
	if err != nil {
		zap.L().Error("error update user total earned value", zap.Error(err))
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
		fmt.Println(err)
		zap.L().Error("error get user by tg id", zap.Error(err))
		return api.NewGetUserProgressBadRequest()
	}

	referralBalance, err := h.UserReferralBalanceCount(ctx, user.ID)
	if err != nil {
		fmt.Println(err)
		zap.L().Error("error count referral balance", zap.Error(err))
		return api.NewGetUserProgressInternalServerError()
	}

	return api.NewGetUserProgressOK().WithPayload(&definition.User{
		TgID:            &user.TgID.Int64,
		Balance:         &user.Balance,
		Wallet:          &user.Wallet.String,
		Energy:          &user.Energy,
		Username:        &user.Username.String,
		ReferralBalance: referralBalance,
	})
}

func (h *Handler) BindUserWalletHandler(req api.BindUserWalletParams) middleware.Responder {
	zap.L().Info(fmt.Sprintf("bind wallet request, user tg id: %d", req.TgID))
	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("error get user by tg id", zap.Error(err))
		return api.NewBindUserWalletBadRequest()
	}

	err = h.userUsecase.UpdateWalletByTgID(ctx, *req.WalletBody.Wallet, user.TgID.Int64)
	if err != nil {
		zap.L().Error("error update user wallet", zap.Error(err))
		return api.NewBindUserWalletInternalServerError()
	}

	return api.NewBindUserWalletOK().WithPayload(&definition.Error{
		Message: &model.WalletWasUpdated,
	})
}

func (h *Handler) GetReferralLinkHandler(req api.GetReferralLinkParams) middleware.Responder {
	zap.L().Info("get referral link request")
	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("error get user", zap.Error(err))
		return api.NewGetReferralLinkBadRequest().WithPayload(&definition.Error{
			Message: &model.NoSuchUser,
		})
	}

	referralLink := h.botLink + "?startapp=" + user.ReferralLink

	return api.NewGetReferralLinkOK().WithPayload(&definition.ReferralLink{
		ReferralLink: &referralLink,
	})

}

func (h *Handler) CreateReferralUserHandler(req api.CreateReferralUserParams) middleware.Responder {
	zap.L().Info("create referral user request")
	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("error get user", zap.Error(err))
		return api.NewCreateReferralUserBadRequest().WithPayload(&definition.Error{
			Message: &model.NoSuchUser,
		})
	}

	if user.ReferralUserID.Valid {
		return api.NewCreateReferralUserBadRequest()
	}

	referralUser, err := h.userUsecase.GetUserByReferralLink(ctx, *req.ReferralLink.ReferralLink)
	if err != nil {
		zap.L().Error("error get referral user", zap.Error(err))
		return api.NewCreateReferralUserBadRequest().WithPayload(&definition.Error{
			Message: &model.NoSuchUser,
		})
	}

	if user.TgID.Int64 == referralUser.TgID.Int64 {
		zap.L().Error("user is referral user by himself", zap.Error(err))
		return api.NewCreateReferralUserBadRequest().WithPayload(&definition.Error{
			Message: &model.NoSuchUser,
		})
	}

	err = h.userUsecase.UpdateUserReferralByTgID(ctx, referralUser.ID, user.TgID.Int64)
	if err != nil {
		zap.L().Error("error update referral part of user", zap.Error(err))
		return api.NewCreateReferralUserInternalServerError()
	}

	return api.NewCreateReferralUserOK().WithPayload(&definition.Error{
		Message: &model.UserHaveReferralUser,
	})
}

func (h *Handler) CollectReferralEarnHandler(req api.CollectReferralEarnParams) middleware.Responder {
	zap.L().Info("collect referral earn request")

	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("error get user", zap.Error(err))
		return api.NewCollectReferralEarnBadRequest().WithPayload(&definition.Error{
			Message: &model.NoSuchUser,
		})
	}

	referralEarn, err := h.CollectReferralEarn(ctx, user)
	if err != nil {
		zap.L().Error("error collect user referral earn", zap.Error(err))
		return api.NewCollectReferralEarnInternalServerError()
	}

	return api.NewCollectReferralEarnOK().WithPayload(&definition.ReferralBalance{
		ReferralBalance: &referralEarn,
	})
}

func (h *Handler) StartEnergyCollectHandler(req api.StartEnergyOfflineParams) middleware.Responder {
	zap.L().Info("start energy collect request")
	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("error get user", zap.Error(err))
		return api.NewStartEnergyOfflineBadRequest().WithPayload(&definition.Error{
			Message: &model.NoSuchUser,
		})
	}

	energy, updatedAt, err := jsonwebtoken.ParseEnergyToken(*req.StartEnergyCollectBody.CurrentEnergyToken, h.tokenSecretKey)
	if err != nil {
		zap.L().Error(fmt.Sprintf("error parse energy token, userID: %d", user.ID), zap.Error(err))
		return api.NewStartEnergyOfflineInternalServerError()
	}

	if user.UpdatedAt.Time.Unix() >= updatedAt {
		zap.L().Error("energy token was expired", zap.Error(err))
		return api.NewStartEnergyOfflineBadRequest()
	}

	err = h.userUsecase.UpdateEnergyByTgID(ctx, energy, user.TgID.Int64)
	if err != nil {
		zap.L().Error("error update user energy", zap.Error(err))
		return api.NewStartEnergyOfflineInternalServerError()
	}

	user.Energy = energy
	userCards, err := h.userCardUsecase.GetUserCardsByUserID(ctx, user.ID)
	if err != nil {
		return api.NewStartEnergyOfflineInternalServerError()
	}

	userEnergyCard := model.UserCard{}

	for i := range userCards {
		if userCards[i].CardName == model.EnergyBooster {
			userEnergyCard = userCards[i]
		}
	}

	if cancelFunc, exists := h.energyCollectTasks[user.TgID.Int64]; exists {
		cancelFunc()
	}

	newCtx, cancel := context.WithCancel(context.Background())
	h.energyCollectTasks[user.TgID.Int64] = cancel

	go func() {
		h.startEnergyCollect(newCtx, user, userEnergyCard)
	}()

	return api.NewStartEnergyOfflineOK().WithPayload(&definition.Error{
		Message: &model.StartedEnergyCollect,
	})

}

func (h *Handler) StopEnergyCollectHandler(req api.StopEnergyOfflineParams) middleware.Responder {
	zap.L().Info("stop energy collect request")

	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("error get user by tg id", zap.Error(err))
		return api.NewStopEnergyOfflineBadRequest()
	}

	if cancelFunc, exists := h.energyCollectTasks[user.TgID.Int64]; exists {
		cancelFunc()
		delete(h.energyCollectTasks, user.TgID.Int64)
	}

	return api.NewStopEnergyOfflineOK().WithPayload(&definition.Error{
		Message: &model.StoppedEnergyCollect,
	})
}

func (h *Handler) startEnergyCollect(ctx context.Context, user model.User, userEnergyCard model.UserCard) {
	zap.L().Info("starting background energy collection", zap.Int64("tgID", user.TgID.Int64))

	baseEnergy := h.EnergyBaseCapacity
	energyUpgrade := h.EnergyUpgrade
	ticker := time.NewTicker(h.EnergyTickTime)
	saveTicker := time.NewTicker(h.EnergySaveTicksTime)
	defer ticker.Stop()
	defer saveTicker.Stop()

	var maxEnergy int64
	if userEnergyCard.CurrentLevel == 1 {
		maxEnergy = int64(baseEnergy)
	} else {
		maxEnergy = int64(baseEnergy) + userEnergyCard.CurrentLevel*int64(energyUpgrade)
	}
	currentEnergy := user.Energy
	fmt.Println(maxEnergy)

	for {
		select {
		case <-ticker.C:
			if currentEnergy < maxEnergy {
				currentEnergy++
				zap.L().Info("energy incremented", zap.Int64("tgID", user.TgID.Int64), zap.Int("energy", int(currentEnergy)))
			}

		case <-saveTicker.C:
			saveCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			err := h.userUsecase.UpdateEnergyByTgID(saveCtx, currentEnergy, user.TgID.Int64)
			cancel()

			if err != nil {
				zap.L().Error("failed to save energy", zap.Error(err))
			} else {
				zap.L().Info("energy saved", zap.Int64("tgID", user.TgID.Int64), zap.Int("energy", int(currentEnergy)))
			}

		case <-ctx.Done():
			// Final save on exit
			saveCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			err := h.userUsecase.UpdateEnergyByTgID(saveCtx, currentEnergy, user.TgID.Int64)
			cancel()

			if err != nil {
				zap.L().Error("failed to save energy on exit", zap.Error(err))
			}
			zap.L().Info("energy collection stopped", zap.Int64("tgID", user.TgID.Int64))
			return
		}
	}
}

func (h *Handler) UserReferralBalanceCount(ctx context.Context, userID int64) (int64, error) {
	zap.L().Info("count user referral balance request")

	var userReferralBalance int64
	var referralsDuty float64
	var referralsSentToUser int64
	referralUsers, err := h.userUsecase.GetUsersByReferralUser(ctx, userID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	for i := range referralUsers {
		referralsDuty += float64(referralUsers[i].TotalCoinEarned) * h.referralUserPart
		referralsSentToUser += referralUsers[i].TotalCoinSentToReferral
	}

	userReferralBalance = int64(referralsDuty) - referralsSentToUser
	return userReferralBalance, nil
}

func (h *Handler) CollectReferralEarn(ctx context.Context, user model.User) (int64, error) {
	zap.L().Info("collect user referral balance request")

	var referralEarn int64
	referralUsers, err := h.userUsecase.GetUsersByReferralUser(ctx, user.ID)
	if err != nil {
		return 0, err
	}

	updateValue := user.Balance

	for i := range referralUsers {
		referralTotalDuty := float64(referralUsers[i].TotalCoinEarned) * h.referralUserPart
		referralAlreadySent := referralUsers[i].TotalCoinSentToReferral
		userEarn := int64(referralTotalDuty) - referralAlreadySent
		updateValue += userEarn

		if referralUsers[i].Balance >= userEarn {
			err = h.userUsecase.UpdateUserBalance(ctx, referralUsers[i].TgID.Int64, referralUsers[i].Balance-userEarn)
			if err != nil {
				return 0, err
			}

			err = h.userUsecase.UpdateTotalCoinSentToReferralByTgID(ctx, referralAlreadySent+userEarn, referralUsers[i].TgID.Int64)
			if err != nil {
				return 0, err
			}

		}

		referralEarn += userEarn
	}

	err = h.userUsecase.UpdateUserBalance(ctx, user.TgID.Int64, updateValue)
	if err != nil {
		return 0, err
	}

	return referralEarn, nil
}
