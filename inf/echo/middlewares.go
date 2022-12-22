package echo

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"userSL/inf/pgsql"
	"userSL/models"
)

func checkLastAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, _ := c.Get("db").(pgsql.Storage)
		login := c.Param("login")
		user, err := db.Load(login)
		if err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		if user.Rule == models.Admin {
			// LastAdmin Запрос количества админов, если один то True
			if db.LastAdmin() {
				log.Println("Attempt to remove the last admin")
				return c.JSON(http.StatusBadRequest, models.JSONResult{"Do not delete. This is the last admin!"})
			}

		}
		// Передаю юзера, чтобы заново не запрашивать.
		c.Set("user", &user)
		return next(c)
	}
}
func forAdmin(login, pass string, c echo.Context) (bool, error) {
	db, _ := c.Get("db").(pgsql.Storage)
	u, err := db.Load(login)
	if err == nil && u.Password == pass {

		if u.Rule == models.Admin {
			return true, nil
		} else {
			log.Println("No access rights ", login)
			c.String(http.StatusForbidden, "forbidden")
		}
	}
	return false, err
}

func forAll(login, pass string, c echo.Context) (bool, error) {
	db, _ := c.Get("db").(pgsql.Storage)
	u, err := db.Load(login)
	if err == nil && u.Password == pass {

		if u.Rule == models.Lock {
			log.Println("No access rights ", login)
			c.String(http.StatusForbidden, "forbidden")
			return false, err
		} else {
			return true, nil
		}
	}
	return false, err
}
