package handler

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/pskoob/miniappBack/internal/api/server/restapi/api"
)

func (h *Handler) SaveUserProgressHandler(req api.SaveProgressParams) middleware.Responder {
	return api.NewSaveProgressOK()
}
