package entities

type Device struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Resolution  Resolution  `json:"resolution"`
	Orientation Orientation `json:"orientation"`
	StatusCode  int8        `json:"statusCode"`
}

type Resolution struct {
	Width  int
	Height int
}

type Orientation int

const (
	OrientationPortrait  Orientation = 0
	OrientationLandscape Orientation = 1
)
