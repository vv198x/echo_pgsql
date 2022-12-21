package pgsql

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"userSL/models"
)

type pgSQL struct {
	c *pg.DB
}

func GetPostgre() *pgSQL {
	options := &pg.Options{
		User:     "pgsql",
		Password: "BAGFIxtgR2ynIL3tip4a",
		Addr:     "localhost:5432",
		Database: "pgdb",
		PoolSize: 50,
	}

	con := pg.Connect(options)
	if con == nil {
		log.Fatal("cannot connect to postgres")
	}

	//con.AddQueryHook(dbLogger{})

	return &pgSQL{con}
}

func (pg *pgSQL) Load(login string) (models.User, error) {
	user := new(models.User)
	err := pg.c.Model(user).Where("login = ?0", login).Select()
	return *user, err
}

func (pg *pgSQL) LoadAll() (*[]models.User, error) {
	users := new([]models.User)
	err := pg.c.Model(users).Select()

	return users, err
}

func (pg *pgSQL) Save(user *models.User) error {
	_, err := pg.c.Model(user).Returning("*").Insert()
	return err
}
func (pg *pgSQL) Change(oldLogin string, user *models.User) error {

	_, err := pg.c.Model(user).Where("login = ?0", oldLogin).Returning("*").Update()

	return err
}
func (pg *pgSQL) Remove(login string) error {
	_, err := pg.c.Model(&models.User{}).Where("login = ?0", login).Delete()
	return err
}

func (pg *pgSQL) CloseDB() error {
	pg.c.Close()
	return nil
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {
	logPG := q.Result
	log.Printf("%v \n", logPG)
}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	logPG, _ := q.FormattedQuery()
	log.Printf("%v \n", logPG)
}

func ReplaceTable(file string) {
	db := GetPostgre()
	defer db.CloseDB()

	c, ioErr := ioutil.ReadFile("./table.sql")
	if ioErr == nil {
		_, err := db.c.Exec(string(c))
		if err != nil {
			fmt.Println(err)
		}
	}

}
