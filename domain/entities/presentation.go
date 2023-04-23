package entities

import "davinci/util"

type Presentation struct {
	Id         int64         `json:"id"`
	Name       string        `json:"name"`
	Resolution Resolution    `json:"resolution"`
	Pages      []Page        `json:"pages,omitempty"`
	StatusCode int64         `json:"statusCode"`
	CreatedAt  util.DateTime `json:"createdAt"`
	ModifiedAt util.DateTime `json:"modifiedAt"`
}

type Page struct {
	Id             int64         `json:"id"`
	IdPresentation int64         `json:"idPresentation"`
	Timing         int64         `json:"timing"`
	Metadata       interface{}   `json:"metadata"`
	StatusCode     int64         `json:"statusCode"`
	CreatedAt      util.DateTime `json:"createdAt"`
	ModifiedAt     util.DateTime `json:"modifiedAt"`
}
