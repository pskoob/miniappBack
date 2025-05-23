// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// BindUserWalletHandlerFunc turns a function with the right signature into a bind user wallet handler
type BindUserWalletHandlerFunc func(BindUserWalletParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BindUserWalletHandlerFunc) Handle(params BindUserWalletParams) middleware.Responder {
	return fn(params)
}

// BindUserWalletHandler interface for that can handle valid bind user wallet params
type BindUserWalletHandler interface {
	Handle(BindUserWalletParams) middleware.Responder
}

// NewBindUserWallet creates a new http.Handler for the bind user wallet operation
func NewBindUserWallet(ctx *middleware.Context, handler BindUserWalletHandler) *BindUserWallet {
	return &BindUserWallet{Context: ctx, Handler: handler}
}

/*
	BindUserWallet swagger:route POST /bind_user_wallet/{tg_id} User bindUserWallet

Bind User Wallet
*/
type BindUserWallet struct {
	Context *middleware.Context
	Handler BindUserWalletHandler
}

func (o *BindUserWallet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewBindUserWalletParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
