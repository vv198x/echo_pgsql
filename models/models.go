package models

type User struct {
	tableName interface{} `pg:"users"`
	Login     string      `json:"login" validate:"required" pg:"login"`
	Password  string      `json:"password" validate:"required" pg:"password"`
	Rule      int         `json:"rule" validate:"gte=0,lt=3" pg:"rule"`
	Name      string      `json:"name" validate:"required" pg:"name"`
	LastName  string      `json:"last_name" validate:"required" pg:"last_name"`
	Dob       int64       `json:"dob" validate:"required" pg:"dob"`
}

const (
	Admin = iota
	Read
	Lock
)
