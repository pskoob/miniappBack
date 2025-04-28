package handler

import (
	// "context"
	// "fmt"

	"context"
	"net/http"
	"strings"
	"time"

	"github.com/pskoob/miniappBack/internal/api/server/restapi"
	"github.com/pskoob/miniappBack/internal/api/server/restapi/api"
	"github.com/pskoob/miniappBack/model"

	"encoding/json"

	"github.com/go-openapi/loads"
	"go.uber.org/zap"
)

type Handler struct {
	userUsecase     model.IUserUsecase
	cardUsecase     model.ICardUsecase
	userCardUsecase model.IUserCardUsecase
	costCardUsecase model.ICostCardUsecase

	autoClickerTasks   map[int64]context.CancelFunc
	energyCollectTasks map[int64]context.CancelFunc

	secretKey        string
	sendingAccount   string
	tokenSecretKey   string
	botLink          string
	hashSalt         string
	referralUserPart float64

	AutoClickerWorkTime      time.Duration
	AutoClickerTickTime      time.Duration
	AutoClickerSaveTicksTime time.Duration

	EnergyBaseCapacity  int64
	EnergyUpgrade       int64
	EnergyTickTime      time.Duration
	EnergySaveTicksTime time.Duration

	router http.Handler
}

func New(
	userUsecase model.IUserUsecase,
	cardUsecase model.ICardUsecase,
	userCardUsecase model.IUserCardUsecase,
	costCardUsecase model.ICostCardUsecase,

	autoClickerTasks map[int64]context.CancelFunc,
	energyCollectTasks map[int64]context.CancelFunc,

	secretKey string,
	sendingAccount string,
	tokenSecretKey string,
	botLink string,
	hashSalt string,
	referralUserPart float64,

	AutoClickerWorkTime time.Duration,
	AutoClickerTickTime time.Duration,
	AutoClickerSaveTicksTime time.Duration,

	EnergyBaseCapacity int64,
	EnergyUpgrade int64,
	EnergyTickTime time.Duration,
	EnergySaveTicksTime time.Duration,
) *Handler {

	withChangedVersion := strings.ReplaceAll(string(restapi.SwaggerJSON), "development", "1")
	swagger, err := loads.Analyzed(json.RawMessage(withChangedVersion), "")
	if err != nil {
		panic(err)
	}

	h := &Handler{
		userUsecase:     userUsecase,
		cardUsecase:     cardUsecase,
		userCardUsecase: userCardUsecase,
		costCardUsecase: costCardUsecase,

		autoClickerTasks:   autoClickerTasks,
		energyCollectTasks: energyCollectTasks,

		secretKey:        secretKey,
		sendingAccount:   sendingAccount,
		tokenSecretKey:   tokenSecretKey,
		botLink:          botLink,
		hashSalt:         hashSalt,
		referralUserPart: referralUserPart,

		AutoClickerWorkTime:      AutoClickerWorkTime,
		AutoClickerTickTime:      AutoClickerTickTime,
		AutoClickerSaveTicksTime: AutoClickerSaveTicksTime,

		EnergyBaseCapacity:  EnergyBaseCapacity,
		EnergyUpgrade:       EnergyUpgrade,
		EnergyTickTime:      EnergyTickTime,
		EnergySaveTicksTime: EnergySaveTicksTime,
	}

	zap.L().Error("server http handler request")
	router := api.NewMaxonBackAPI(swagger)
	router.UseSwaggerUI()
	router.Logger = zap.S().Infof

	// USER
	router.GetUserProgressHandler = api.GetUserProgressHandlerFunc(h.GetUserHandler)
	router.SaveProgressHandler = api.SaveProgressHandlerFunc(h.SaveUserProgressHandler)
	router.BindUserWalletHandler = api.BindUserWalletHandlerFunc(h.BindUserWalletHandler)
	router.GetReferralLinkHandler = api.GetReferralLinkHandlerFunc(h.GetReferralLinkHandler)
	router.CreateReferralUserHandler = api.CreateReferralUserHandlerFunc(h.CreateReferralUserHandler)
	router.CollectReferralEarnHandler = api.CollectReferralEarnHandlerFunc(h.CollectReferralEarnHandler)

	// AUTO CLICKER
	router.StartAutoClickerHandler = api.StartAutoClickerHandlerFunc(h.StartAutoClickerHandler)
	router.StopAutoClickerHandler = api.StopAutoClickerHandlerFunc(h.StopAutoClickerHandler)

	// NEAR
	router.TransitNearHandler = api.TransitNearHandlerFunc(h.NearTransactionHandler)

	// USER CARDS
	router.UpdateUserCardHandler = api.UpdateUserCardHandlerFunc(h.UpdateUserCardHandler)
	router.GetUserCardsHandler = api.GetUserCardsHandlerFunc(h.GetUserCardsHandler)

	// ENERGY
	router.StartEnergyOfflineHandler = api.StartEnergyOfflineHandlerFunc(h.StartEnergyCollectHandler)
	router.StopEnergyOfflineHandler = api.StopEnergyOfflineHandlerFunc(h.StopEnergyCollectHandler)

	h.router = router.Serve(nil)

	return h
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	zap.L().Info("Received HTTP request", zap.String("method", r.Method), zap.String("path", r.URL.Path))

	if h.router == nil {
		zap.L().Error("h.router is nil")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	zap.L().Info("h.router is not nil, processing request")
	h.router.ServeHTTP(w, r)
}
