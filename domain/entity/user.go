package entity

type User struct {
	entity
	UUID     string
	Name     string
	CPF      string
	Email    string
	Password string
}
