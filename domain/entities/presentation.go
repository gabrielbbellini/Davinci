package entities

import "davinci/util"

type Presentation struct {
	Id           int64         `json:"id"`
	Name         string        `json:"name"`
	ResolutionId int64         `json:"resolutionId"`
	Orientation  Orientation   `json:"orientation"`
	Pages        []Page        `json:"pages,omitempty"`
	StatusCode   int64         `json:"statusCode"`
	CreatedAt    util.DateTime `json:"createdAt"`
	ModifiedAt   util.DateTime `json:"modifiedAt"`
}

type Page struct {
	Id             int64         `json:"id"`
	PresentationId int64         `json:"presentationId"`
	Duration       int64         `json:"duration"`
	Component      interface{}   `json:"component"`
	StatusCode     int64         `json:"statusCode"`
	CreatedAt      util.DateTime `json:"createdAt"`
	ModifiedAt     util.DateTime `json:"modifiedAt"`
}
