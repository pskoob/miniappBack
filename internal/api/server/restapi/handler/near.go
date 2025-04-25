package handler

import (
	"math/big"

	"github.com/go-openapi/runtime/middleware"
	"github.com/pskoob/miniappBack/internal/api/definition"
	"github.com/pskoob/miniappBack/internal/api/server/restapi/api"
	"github.com/pskoob/miniappBack/model"
	"github.com/pskoob/miniappBack/pkg/nearffi"

	"go.uber.org/zap"
)

type FTResult struct {
	Status  string
	Message string
}

func (h *Handler) NearTransactionHandler(req api.TransitNearParams) middleware.Responder {
	zap.L().Info("near transit request")
	//ctx := req.HTTPRequest.Context()

	receivingAccount := req.GetNearBody.ReceivingWallet

	amountNear := req.GetNearBody.Amount
	yocto := new(big.Int)
	yocto.Mul(big.NewInt(int64(amountNear*1e6)), big.NewInt(1e18))
	amountStr := yocto.String()

	result, err := nearffi.SendFT(h.sendingAccount, receivingAccount, amountStr, h.secretKey, req.GetNearBody.TokenContract)
	if err != nil {
		zap.L().Error("error send money", zap.Error(err))
		return api.NewTransitNearBadRequest().WithPayload(&definition.Error{
			Message: &model.CryptoDidNotSend,
		})
	}

	return api.NewTransitNearOK().WithPayload(&definition.NearTransit{
		Result: result.Message,
	})
}
