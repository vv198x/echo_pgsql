package pgsql

import (
	"github.com/go-pg/pg"
	"github.com/labstack/gommon/log"
	"userSL/models"
	"userSL/pkg/cfg"
)

type pgSQL struct {
	c *pg.DB
}

func GetPostgre() *pgSQL {
	con := pg.Connect(&pg.Options{
		User:     cfg.Get().PGUser,
		Password: cfg.Get().PGPass,
		Addr:     cfg.Get().PGAddr,
		Database: cfg.Get().PGDB,
		PoolSize: 50,
	})
	if con == nil {
		log.Fatal("cannot connect to postgres")
	}

	//Для захвата и логирования запросов
	if cfg.Get().Debug {
		con.AddQueryHook(dbLogger{})
	}

	return &pgSQL{con}
}

func (pg *pgSQL) Load(login string) (models.User, error) {
	u := getDB(&models.User{})

	err := pg.c.Model(u).Where("login = ?0", login).Select()
	return u.convUser(), err
}

func (pg *pgSQL) LoadAll() ([]models.User, error) {
	usersDB := []userDB{}

	err := pg.c.Model(&usersDB).Select()

	users := []models.User{}
	for _, u := range usersDB {
		users = append(users, u.convUser())
	}

	return users, err
}

func (pg *pgSQL) Save(user *models.User) error {
	u := getDB(user)

	_, err := pg.c.Model(u).OnConflict("(login) DO NOTHING").Returning("*").Insert()
	//Вернуть результат Returning
	*user = u.convUser()
	return err
}
func (pg *pgSQL) Change(oldLogin string, user *models.User) error {
	u := getDB(user)

	_, err := pg.c.Model(u).Where("login = ?0", oldLogin).Returning("*").Update()
	*user = u.convUser()
	return err
}

func (pg *pgSQL) LastAdmin() (last bool) {
	//Запрос вернет 1 если есть только один администратор
	pg.c.Query(&last, `select cast(case when COUNT(*) > 1 then 0 else 1 end AS BIT)
							from users 
							where rule = ?0`, models.Admin)
	return
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
	_, err := pg.c.Model(&userDB{}).Where("login = ?0", login).Delete()
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
