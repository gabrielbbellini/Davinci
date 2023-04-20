package entities

type Orientation string

const (
	OrientationPortrait  Orientation = "portrait"
	OrientationLandscape Orientation = "landscape"
)

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
