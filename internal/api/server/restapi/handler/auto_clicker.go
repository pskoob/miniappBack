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

	if user.AutoClicker {
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
	duration := 3 * time.Hour                 // Время работы автокликера
	ticker := time.NewTicker(1 * time.Second) // Создаем тикер, который будет срабатывать каждую секунду
	defer ticker.Stop()                       // Убедимся, что тикер остановится, когда функция завершится

	endTime := time.Now().Add(duration) // Время окончания работы автокликера
	balance := user.Balance.Int64
	for {
		select {
		case <-ticker.C: // Каждую секунду выполняем действие
			// Здесь вы можете реализовать логику "тапа"
			oneTap := user.ClickBooster.Int64 + 1
			balance += 1
			zap.L().Info("Auto clicker tapped for user", zap.Int64("tgID", user.TgID.Int64), zap.Int64("oneTap", oneTap))
			zap.L().Info("user balance: ", zap.Int64("", balance))

			// Если время работы автокликера истекло, выходим из цикла
			if time.Now().After(endTime) {
				// Создаем новый контекст для сохранения данных
				saveCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()

				err := h.userUsecase.SaveProgressByTgID(saveCtx, user.AutoClicker, user.ClickBooster.Int64, balance, user.TgID.Int64)
				if err != nil {
					zap.L().Error("error save user data", zap.Error(err))
				}

				zap.L().Info("Auto clicker finished for user", zap.Int64("tgID", user.TgID.Int64))
				return
			}

		case <-ctx.Done(): // Если контекст отменен, выходим из цикла
			// Создаем новый контекст для сохранения данных
			saveCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			err := h.userUsecase.SaveProgressByTgID(saveCtx, user.AutoClicker, user.ClickBooster.Int64, balance, user.TgID.Int64)
			if err != nil {
				zap.L().Error("error save user data", zap.Error(err))
			}
			zap.L().Info("Auto clicker stopped for user", zap.Int64("tgID", user.TgID.Int64))
			return
		}
	}
}
