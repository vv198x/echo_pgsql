package echo

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"strings"
	"userSL/inf/pgsql"
	"userSL/models"
)

func headlers(e *echo.Echo) {
	root := e.Group("/api/users/v1", validJSON)
	login := e.Group("/api/users/v1/:login", validJSON, checkLogin)

	root.POST("/", Create, middleware.BasicAuth(forAdmin))
	login.PUT("", Update, middleware.BasicAuth(forAdmin))
	login.DELETE("", Delete, checkLastAdmin, middleware.BasicAuth(forAdmin))
	// Для пакета swag разделил Read(на Read и ReadAll). И не получилось в одну совместить
	// На верно нужен костомный хендлер
	root.GET("/", ReadAll, middleware.BasicAuth(forAll))
	login.GET("", Read, middleware.BasicAuth(forAll))

}

func validJSON(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)
		if c.Request().Method != http.MethodGet {

			c.Bind(user)
			err := c.Validate(user)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, err.Error())
			}
			// Передаю валидного юзера дальше
			c.Set("user", &user)
		}

		return next(c)
	}
}

func checkLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, _ := c.Get("db").(pgsql.Storage)
		login := c.Param("login")
		user, err := db.Load(login)
		if err != nil {

			if strings.Contains(err.Error(), "no rows") {
				return c.JSON(http.StatusNotFound, models.JSONResult{"Not found"})
			} else {

				log.Println("DB error", err)
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}

		// Передаю юзера дальше
		c.Set("user", &user)
		return next(c)
	}
}

func checkLastAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, _ := c.Get("db").(pgsql.Storage)
		user := *(c.Get("user").(*models.User))

		if user.Rule == models.Admin {
			// LastAdmin Запрос количества админов, если один то True
			if db.LastAdmin() {
				log.Println("Attempt to remove the last admin")
				return c.JSON(http.StatusBadRequest, models.JSONResult{"Do not delete. This is the last admin!"})
			}

		}

		return next(c)
	}
}
