// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ColectReferralEarnHandlerFunc turns a function with the right signature into a colect referral earn handler
type ColectReferralEarnHandlerFunc func(ColectReferralEarnParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ColectReferralEarnHandlerFunc) Handle(params ColectReferralEarnParams) middleware.Responder {
	return fn(params)
}

// ColectReferralEarnHandler interface for that can handle valid colect referral earn params
type ColectReferralEarnHandler interface {
	Handle(ColectReferralEarnParams) middleware.Responder
}

// NewColectReferralEarn creates a new http.Handler for the colect referral earn operation
func NewColectReferralEarn(ctx *middleware.Context, handler ColectReferralEarnHandler) *ColectReferralEarn {
	return &ColectReferralEarn{Context: ctx, Handler: handler}
}

/*
	ColectReferralEarn swagger:route GET /collect_referral_earn/{tg_id} User colectReferralEarn

Collect Referral Earn
*/
type ColectReferralEarn struct {
	Context *middleware.Context
	Handler ColectReferralEarnHandler
}

func (o *ColectReferralEarn) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewColectReferralEarnParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
