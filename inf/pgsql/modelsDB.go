package pgsql

import (
	"log"
	"time"
	"userSL/models"
)

type userDB struct {
	tableName interface{} `sql:"users"`
	Login     string      `sql:"login"`
	Password  string      `sql:"password"`
	Rule      int         `sql:"rule"`
	Name      string      `sql:"name"`
	LastName  string      `sql:"last_name"`
	Dob       time.Time   `sql:"dob"`
}

func (uDB *userDB) convUser() models.User {
	return models.User{
		Login:    uDB.Login,
		Password: uDB.Password,
		Rule:     uDB.Rule,
		Name:     uDB.Name,
		LastName: uDB.LastName,
		Dob:      uDB.Dob.Format("02-01-2006"),
	}
}

func getDB(u *models.User) *userDB {
	timestamp, err := time.Parse("02-01-2006", u.Dob)
	if err != nil {
		log.Println("Wrong dob ", err)
	}
	return &userDB{
		Login:    u.Login,
		Password: u.Password,
		Rule:     u.Rule,
		Name:     u.Name,
		LastName: u.LastName,
		Dob:      timestamp}
}
