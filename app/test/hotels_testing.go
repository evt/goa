// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "hotels": hotels TestHelpers
//
// Command:
// $ goagen
// --design=goa/design
// --out=$(GOPATH)/src/hotels
// --version=v1.3.0

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"goa/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
)

// SearchHotelsOK runs the method Search of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func SearchHotelsOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.HotelsController) (http.ResponseWriter, app.HotelCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/hotel/search"),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "HotelsTest"), rw, req, prms)
	searchCtx, _err := app.NewSearchHotelsContext(goaCtx, req, service)
	if _err != nil {
		panic("invalid test data " + _err.Error()) // bug
	}

	// Perform action
	_err = ctrl.Search(searchCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.HotelCollection
	if resp != nil {
		var ok bool
		mt, ok = resp.(app.HotelCollection)
		if !ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.HotelCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}
