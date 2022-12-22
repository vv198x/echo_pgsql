package pgsql

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.com/labstack/gommon/log"
	"io/ioutil"
	"userSL/models"
	"userSL/pkg/config"
)

type pgSQL struct {
	c *pg.DB
}

func GetPostgre() *pgSQL {
	con := pg.Connect(&pg.Options{
		User:     *config.PGUser,
		Password: *config.PGPass,
		Addr:     *config.PGAddr,
		Database: *config.PGDB,
		PoolSize: *config.PGPoolSize,
	})
	if con == nil {
		log.Fatal("cannot connect to postgres")
	}

	//Для захвата и логирования запросов
	if *config.Debug {
		con.AddQueryHook(dbLogger{})
	}

	return &pgSQL{con}
}

func (pg *pgSQL) Load(login string) (models.User, error) {
	user := new(models.User)

	err := pg.c.Model(user).Where("login = ?0", login).Select()
	return *user, err
}

func (pg *pgSQL) LoadAll() ([]models.User, error) {
	users := new([]models.User)

	err := pg.c.Model(users).Select()
	return *users, err
}

func (pg *pgSQL) Save(user *models.User) error {

	_, err := pg.c.Model(user).Returning("*").Insert()
	return err
}
func (pg *pgSQL) Change(oldLogin string, user *models.User) error {

	_, err := pg.c.Model(user).Where("login = ?0", oldLogin).Returning("*").Update()
	return err
}

func (pg *pgSQL) LastAdmin() bool {
	var count int
	_, err := pg.c.Query(&count, `select COUNT(*) from users where rule = 0`)
	if err == nil && count == 1 {
		return true
	}
	return false
}
func (pg *pgSQL) Remove(login string, rule int) error {
	if rule == models.Admin {
		//Запрос удалить админа, выполнится если админы ещё есть.
		_, err := pg.c.Exec(`
		delete
		from users
		where login = ?0 and 1 < (select count(*)
			from users
			where rule = ?1)
			`, login, models.Admin)

		return err

	}
	_, err := pg.c.Model(&models.User{}).Where("login = ?0", login).Delete()
	return err
}

func (pg *pgSQL) CloseDB() error {
	pg.c.Close()
	return nil
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(q *pg.QueryEvent) {

}

func (d dbLogger) AfterQuery(q *pg.QueryEvent) {
	logPG, _ := q.FormattedQuery()
	log.Printf("%v \n", logPG)
}

func ReplaceTable(file string) {
	db := GetPostgre()
	defer db.CloseDB()

	c, ioErr := ioutil.ReadFile(file)
	if ioErr == nil {
		_, err := db.c.Exec(string(c))
		if err != nil {
			fmt.Println(err)
		}
	}

}
