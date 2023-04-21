package entities

type Device struct {
	Id          int64       `json:"id"`
	Name        string      `json:"name"`
	Resolution  Resolution  `json:"resolution"`
	Orientation Orientation `json:"orientation"`
}

type Resolution struct {
	Width  int
	Height int
}

type Orientation int

const (
	OrientationPortrait  Orientation = 1
	OrientationLandscape Orientation = 2
)
