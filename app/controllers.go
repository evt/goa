// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "hotels": Application Controllers
//
// Command:
// $ goagen
// --design=goa/design
// --out=$(GOPATH)/src/hotels
// --version=v1.3.0

package app

import (
	"context"
	"github.com/goadesign/goa"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// HotelsController is the controller interface for the Hotels actions.
type HotelsController interface {
	goa.Muxer
	Search(*SearchHotelsContext) error
}

// MountHotelsController "mounts" a Hotels resource controller on the given service.
func MountHotelsController(service *goa.Service, ctrl HotelsController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSearchHotelsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Search(rctx)
	}
	service.Mux.Handle("GET", "/hotel/search", ctrl.MuxHandler("search", h, nil))
	service.LogInfo("mount", "ctrl", "Hotels", "action", "Search", "route", "GET /hotel/search")
}
