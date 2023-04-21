package entities

import "base/util"

type User struct {
	Id         int64         `json:"id"`
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Role       string        `json:"role"`
	PassWord   string        `json:"password"`
	StatusCode int8          `json:"statusCode"`
	ModifiedAt util.DateTime `json:"modifiedAt"`
	CreatedAt  util.DateTime `json:"createdAt"`
}
