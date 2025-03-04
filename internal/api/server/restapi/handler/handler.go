package handler

import (
	// "context"
	// "fmt"

	"context"
	"net/http"
	"strings"

	"github.com/pskoob/miniappBack/internal/api/server/restapi"
	"github.com/pskoob/miniappBack/internal/api/server/restapi/api"
	"github.com/pskoob/miniappBack/model"

	"encoding/json"

	"github.com/go-openapi/loads"
	"go.uber.org/zap"
)

type Handler struct {
	userUsecase      model.IUserUsecase
	autoClickerTasks map[int64]context.CancelFunc

	router http.Handler
}

func New(
	userUsecase model.IUserUsecase,
	autoClickerTasks map[int64]context.CancelFunc,
) *Handler {

	withChangedVersion := strings.ReplaceAll(string(restapi.SwaggerJSON), "development", "1")
	swagger, err := loads.Analyzed(json.RawMessage(withChangedVersion), "")
	if err != nil {
		panic(err)
	}

	h := &Handler{
		userUsecase:      userUsecase,
		autoClickerTasks: autoClickerTasks,
	}

	zap.L().Error("server http handler request")
	router := api.NewMaxonBackAPI(swagger)
	router.UseSwaggerUI()
	router.Logger = zap.S().Infof

	// Progress
	router.GetUserProgressHandler = api.GetUserProgressHandlerFunc(h.GetUserProgressHandler)
	router.SaveProgressHandler = api.SaveProgressHandlerFunc(h.SaveUserProgressHandler)

	// Auto clicker
	router.StartAutoClickerHandler = api.StartAutoClickerHandlerFunc(h.StartAutoClickerHandler)
	router.StopAutoClickerHandler = api.StopAutoClickerHandlerFunc(h.StopAutoClickerHandler)

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
