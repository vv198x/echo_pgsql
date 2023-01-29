package pgsql

import (
	"errors"
	"userSL/models"
)

type mockDB struct {
}

var (
	admin = models.User{Login: "admin", Password: "admin", Rule: models.Admin, Name: "first name Admin", LastName: "last name Admin", Dob: "01-01-1970"}
	user  = models.User{Login: "user", Password: "user", Rule: models.Read, Name: "first name user", LastName: "last name user", Dob: "21-10-2000"}
)

func GetMockDB() *mockDB {
	return &mockDB{}
}

func (*mockDB) Load(login string) (models.User, error) {

	if login == "admin" {
		return admin, nil
	}
	if login == "user" {
		return user, nil
	}
	return models.User{}, errors.New("user not found")
}

func (*mockDB) LoadAll() ([]models.User, error) {
	users := []models.User{admin, user}

	return users, nil
}

func (*mockDB) Save(user *models.User) error {
	return nil
}
func (*mockDB) Change(oldLogin string, user *models.User) error {
	if oldLogin == "admin" || oldLogin == "user" {
		return nil
	}
	return errors.New("user not found")
}
func (*mockDB) Remove(login string, rule int) error {
	if login == "admin" || login == "user" {
		return nil
	}
	return errors.New("user not found")
}

func (*mockDB) LastAdmin() bool {
	return true
}

func (*mockDB) CloseDB() error {
	return nil
}
