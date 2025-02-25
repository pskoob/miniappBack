package handler

import (
	// "context"
	// "fmt"

	"net/http"
	// "strings"


	// "encoding/json"

	"go.uber.org/zap"
	// "github.com/go-openapi/loads"
)

type Handler struct {
	router http.Handler
}

func New() *Handler {

	// withChangedVersion := strings.ReplaceAll(string(restapi.SwaggerJSON), "development", "1")
	// swagger, err := loads.Analyzed(json.RawMessage(withChangedVersion), "")
	// if err != nil {
	// 	panic(err)
	// }

	h := &Handler{}

	zap.L().Error("server http handler request")
	// router := api.NewSpaceVPXBackendServiceAPI(swagger)
	// router.UseSwaggerUI()
	// router.Logger = zap.S().Infof
	

	// h.router = router.Serve(nil)

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
