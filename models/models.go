package models

type User struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
	Rule     int    `json:"rule" validate:"gte=0,lt=3"`
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Dob      string `json:"dob" validate:"required"`
}

const (
	Admin = iota
	Read
	Lock
)
