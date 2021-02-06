// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// PostTestHandlerFunc turns a function with the right signature into a post test handler
type PostTestHandlerFunc func(PostTestParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostTestHandlerFunc) Handle(params PostTestParams) middleware.Responder {
	return fn(params)
}

// PostTestHandler interface for that can handle valid post test params
type PostTestHandler interface {
	Handle(PostTestParams) middleware.Responder
}

// NewPostTest creates a new http.Handler for the post test operation
func NewPostTest(ctx *middleware.Context, handler PostTestHandler) *PostTest {
	return &PostTest{Context: ctx, Handler: handler}
}

/* PostTest swagger:route POST /test postTest

PostTest post test API

*/
type PostTest struct {
	Context *middleware.Context
	Handler PostTestHandler
}

func (o *PostTest) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostTestParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}