package entities

import "base/util"

type User struct {
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	Credential Credential    `json:"credential"`
	CreatedAt  util.DateTime `json:"createdAt"`
	ModifiedAt util.DateTime `json:"modifiedAt"`
	StatusCode string        `json:"statusCode"`
}
