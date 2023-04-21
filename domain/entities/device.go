package entities

import "base/util"

type Device struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	Resolution  Resolution    `json:"resolution"`
	Orientation Orientation   `json:"orientation"`
	StatusCode  int8          `json:"statusCode"`
	ModifiedAt  util.DateTime `json:"modifiedAt"`
	CreatedAt   util.DateTime `json:"createdAt"`
}

type Resolution struct {
	Width      int64         `json:"width"`
	Height     int64         `json:"height"`
	StatusCode int8          `json:"statusCode"`
	ModifiedAt util.DateTime `json:"modifiedAt"`
	CreatedAt  util.DateTime `json:"createdAt"`
}

type Orientation int

const (
	OrientationPortrait  Orientation = 0
	OrientationLandscape Orientation = 1
)
