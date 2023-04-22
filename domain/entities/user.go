package entities

import "base/util"

const StatusOk = 1
const StatusDeleted = 2

type User struct {
	Id         int64         `json:"id"`
	Name       string        `json:"name"`
	Credential Credential    `json:"credential"`
	CreatedAt  util.DateTime `json:"createdAt"`
	ModifiedAt util.DateTime `json:"modifiedAt"`
	StatusCode int64         `json:"statusCode"`
}
