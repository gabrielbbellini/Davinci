package entities

import "base/util"

type Device struct {
	Id          int64         `json:"id,omitempty"`
	Name        string        `json:"name,omitempty"`
	Resolution  *Resolution   `json:"resolution,omitempty"`
	Orientation Orientation   `json:"orientation,omitempty"`
	StatusCode  int8          `json:"statusCode"`
	ModifiedAt  util.DateTime `json:"modifiedAt,omitempty"`
	CreatedAt   util.DateTime `json:"createdAt,omitempty"`
}

type Resolution struct {
	Id         int64         `json:"id,omitempty"`
	Width      int64         `json:"width,omitempty"`
	Height     int64         `json:"height,omitempty"`
	StatusCode int8          `json:"statusCode,omitempty"`
	ModifiedAt util.DateTime `json:"modifiedAt,omitempty"`
	CreatedAt  util.DateTime `json:"createdAt,omitempty"`
}

type Orientation int

const (
	OrientationPortrait  Orientation = 0
	OrientationLandscape Orientation = 1
)
