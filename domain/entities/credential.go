package entities

type Credential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleId   int64  `json:"roleId"`
}
