// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "hotels": Application Media Types
//
// Command:
// $ goagen
// --design=goa/design
// --out=$(GOPATH)/src/hotels
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// Hotel media type (default view)
//
// Identifier: hotel; view=default
type Hotel struct {
	ID   int    `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Hotel media type instance.
func (mt *Hotel) Validate() (err error) {

	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.ID < 0 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.id`, mt.ID, 0, true))
	}
	return
}

// DecodeHotel decodes the Hotel instance encoded in resp body.
func (c *Client) DecodeHotel(resp *http.Response) (*Hotel, error) {
	var decoded Hotel
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// HotelCollection is the media type for an array of Hotel (default view)
//
// Identifier: hotel; type=collection; view=default
type HotelCollection []*Hotel

// Validate validates the HotelCollection media type instance.
func (mt HotelCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeHotelCollection decodes the HotelCollection instance encoded in resp body.
func (c *Client) DecodeHotelCollection(resp *http.Response) (HotelCollection, error) {
	var decoded HotelCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}
