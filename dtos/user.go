package dtos

type User struct {
	UserId int    `json:"user_id"`
	Roles  []Role `json:"roles"`
	Email  string `json:"email"`
}

type Role struct {
	Id   int    `json:"Id"`
	Name string `json:"Role"`
}
