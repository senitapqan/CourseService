package dtos

type User struct {
	UserId int
	Roles  []Role
	Email  string
}

type Role struct {
	Id   int
	Name string
}
