package echo

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
	"userSL/inf/pgsql"
	"userSL/models"
)

func headlers(e *echo.Echo) {

	e.POST("/api/users/v1/auth/", getToken)

	admin := e.Group("/api/users/v1", checkToken, forAdmin)
	read := e.Group("/api/users/v1", checkToken, forAll)

	admin.POST("/", Create, validJSON)
	admin.PUT("/:login", Update, checkLogin, validJSON)
	admin.DELETE("/:login", Delete, checkLogin, checkLastAdmin)

	// Для пакета swag разделил Read(на Read и ReadAll).
	read.GET("/", ReadAll)
	read.GET("/:login", Read, checkLogin)

}

func validJSON(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(models.User)

		err := c.Bind(user) //nolint
		err = c.Validate(user)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// Передаю валидного юзера дальше
		c.Set("validUser", user)

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
				return c.JSON(http.StatusNotFound, JSONResult{"Not found"})
			} else {

				log.Println("DB error", err)
				return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
			}
		}

		// Передаю юзера дальше
		c.Set("userSL", &user)

		return next(c)
	}
}

func checkLastAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db, _ := c.Get("db").(pgsql.Storage)

		user := c.Get("userSL").(*models.User)

		if user.Rule == models.Admin {

			if db.LastAdmin() {
				log.Println("Attempt to remove the last admin")
				return c.JSON(http.StatusBadRequest, JSONResult{"Do not delete. This is the last admin!"})
			}

		}

		return next(c)
	}
}
