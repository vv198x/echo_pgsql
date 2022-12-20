package echo

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"userSL/app/users"
	"userSL/inf/pgsql"
	"userSL/models"
)

func routes(e *echo.Echo) {

	e.POST("/api/users/v1", users.Create, middleware.BasicAuth(forAdmin))
	e.GET("/api/users/v1/", users.Read, middleware.BasicAuth(forAll))
	e.GET("/api/users/v1/:login", users.Read, middleware.BasicAuth(forAll))
	e.PUT("/api/users/v1/:login", users.Update, middleware.BasicAuth(forAdmin))
	e.DELETE("/api/users/v1/:login", users.Delete, middleware.BasicAuth(forAdmin))

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
