// Code generated by go-swagger; DO NOT EDIT.

package api

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// UpdateUserCardHandlerFunc turns a function with the right signature into a update user card handler
type UpdateUserCardHandlerFunc func(UpdateUserCardParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UpdateUserCardHandlerFunc) Handle(params UpdateUserCardParams) middleware.Responder {
	return fn(params)
}

// UpdateUserCardHandler interface for that can handle valid update user card params
type UpdateUserCardHandler interface {
	Handle(UpdateUserCardParams) middleware.Responder
}

// NewUpdateUserCard creates a new http.Handler for the update user card operation
func NewUpdateUserCard(ctx *middleware.Context, handler UpdateUserCardHandler) *UpdateUserCard {
	return &UpdateUserCard{Context: ctx, Handler: handler}
}

/*
	UpdateUserCard swagger:route POST /update_user_card/{tg_id} Cards updateUserCard

Update User Card
*/
type UpdateUserCard struct {
	Context *middleware.Context
	Handler UpdateUserCardHandler
}

func (o *UpdateUserCard) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewUpdateUserCardParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
