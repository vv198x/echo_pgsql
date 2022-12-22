package pgsql

import "userSL/models"

type Storage interface {
	Load(login string) (models.User, error)
	LoadAll() ([]models.User, error)
	Save(user *models.User) error
	Change(oldLogin string, user *models.User) error
	Remove(login string, rule int) error
	LastAdmin() bool
	CloseDB() error
}
