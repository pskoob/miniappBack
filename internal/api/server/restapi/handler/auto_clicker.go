package handler

import (
	"context"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pskoob/miniappBack/internal/api/definition"
	"github.com/pskoob/miniappBack/internal/api/server/restapi/api"
	"github.com/pskoob/miniappBack/model"
	"go.uber.org/zap"
)

func (h *Handler) StartAutoClickerHandler(req api.StartAutoClickerParams) middleware.Responder {
	zap.L().Info("start auto clicker request")

	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("error get user by tg id", zap.Error(err))
		return api.NewGetUserProgressInternalServerError()
	}

	userCards, err := h.userCardUsecase.GetUserCardsByUserID(ctx, user.ID)
	if err != nil {
		zap.L().Error("user do not have boost cards", zap.Error(err))
		return api.NewStartAutoClickerBadRequest().WithPayload(&definition.Error{
			Message: &model.HasNotAutoClicker,
		})
	}

	userAutoClicker := false

	for i := range userCards {
		if userCards[i].AutoClicker.Bool {
			userAutoClicker = true
			break
		}
	}

	if userAutoClicker {
		// Если автокликер уже запущен, отменяем его
		if cancelFunc, exists := h.autoClickerTasks[user.TgID.Int64]; exists {
			cancelFunc() // Останавливаем предыдущий автокликер
		}

		// Создаем новый контекст с возможностью отмены
		newCtx, cancel := context.WithCancel(context.Background()) // Используем контекст Background
		h.autoClickerTasks[user.TgID.Int64] = cancel

		// Запускаем автокликер в горутине
		go func() {
			h.startAutoClickerForUser(newCtx, user)
		}()

		return api.NewStartAutoClickerOK().WithPayload(&definition.Error{
			Message: &model.StartedAutoClicker,
		})
	} else {
		return api.NewStartAutoClickerOK().WithPayload(&definition.Error{
			Message: &model.HasNotAutoClicker,
		})
	}
}

func (h *Handler) StopAutoClickerHandler(req api.StopAutoClickerParams) middleware.Responder {
	zap.L().Info("stop auto clicker request")

	ctx := req.HTTPRequest.Context()

	user, err := h.userUsecase.GetUserByTgID(ctx, req.TgID)
	if err != nil {
		zap.L().Error("error get user by tg id", zap.Error(err))
		return api.NewGetUserProgressInternalServerError()
	}

	if cancelFunc, exists := h.autoClickerTasks[user.TgID.Int64]; exists {
		cancelFunc()
		delete(h.autoClickerTasks, user.TgID.Int64)
	}

	return api.NewStopAutoClickerOK().WithPayload(&definition.Error{
		Message: &model.StoppedAutoClicker,
	})
}

func (h *Handler) startAutoClickerForUser(ctx context.Context, user model.User) {
	duration := 30 * time.Minute                    // Время работы автокликера
	actionTicker := time.NewTicker(1 * time.Second) // Тикер для выполнения действий каждую секунду
	saveTicker := time.NewTicker(5 * time.Second)   // Тикер для сохранения данных каждые n секунд (например, 5 секунд)
	defer actionTicker.Stop()                       // Остановка тикера действий
	defer saveTicker.Stop()                         // Остановка тикера сохранения

	endTime := time.Now().Add(duration) // Время окончания работы автокликера
	balance := user.Balance.Int64

	userCards, err := h.userCardUsecase.GetUserCardsByUserID(ctx, user.ID)
	if err != nil {
		zap.L().Error("error get user cards", zap.Error(err))
	}

	var powerClick model.UserCard

	for {
		select {
		case <-actionTicker.C: // Каждую секунду выполняем действие
			// Логика "тапа"

			for i := range userCards {
				if userCards[i].CardName == model.PowerClick {
					powerClick = userCards[i]
				}
			}

			oneTap := powerClick.CurrentLevel + 1
			balance += 1
			zap.L().Info("Auto clicker tapped for user", zap.Int64("tgID", user.TgID.Int64), zap.Int64("oneTap", oneTap))
			zap.L().Info("user balance: ", zap.Int64("", balance))

		case <-saveTicker.C: // Каждые n секунд сохраняем данные
			saveCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			err := h.userUsecase.UpdateUserBalance(saveCtx, user.TgID.Int64, balance)
			if err != nil {
				zap.L().Error("error save user data", zap.Error(err))
			} else {
				zap.L().Info("User  data saved successfully for tgID", zap.Int64("tgID", user.TgID.Int64))
			}

		case <-ctx.Done(): // Если контекст отменен, выходим из цикла
			// Сохраняем данные перед выходом
			saveCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			err := h.userUsecase.UpdateUserBalance(saveCtx, user.TgID.Int64, balance)
			if err != nil {
				zap.L().Error("error save user data", zap.Error(err))
			}
			zap.L().Info("Auto clicker stopped for user", zap.Int64("tgID", user.TgID.Int64))
			return
		}

		// Если время работы автокликера истекло, выходим из цикла
		if time.Now().After(endTime) {
			// Сохраняем данные перед выходом
			saveCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			err := h.userUsecase.UpdateUserBalance(saveCtx, user.TgID.Int64, balance)
			if err != nil {
				zap.L().Error("error save user data", zap.Error(err))
			}

			zap.L().Info("Auto clicker finished for user", zap.Int64("tgID", user.TgID.Int64))
			return
		}
	}
}
