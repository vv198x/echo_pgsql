package models

type User struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
	Rule     int    `json:"rule" validate:"gte=1,lte=3"` // 1-admin 2-read 3-lock
	Name     string `json:"name" validate:"required"`
	LastName string `json:"last_name" validate:"required"`
	Dob      string `json:"dob" validate:"required"`
}

const (
	Admin = iota + 1
	Read
	Lock
)
