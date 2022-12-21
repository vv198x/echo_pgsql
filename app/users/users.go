package users

import (
	"github.com/labstack/echo"
	"log"
	"net/http"
	"userSL/inf/pgsql"
	"userSL/models"
)

func Read(c echo.Context) error {
	db, _ := c.Get("db").(pgsql.Storage)

	login := c.Param("login")
	if login == "" {
		users, err := db.LoadAll()
		if err != nil {
			log.Println("DB error ", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusFound, &users)
	}

	user, err := db.Load(login)
	if err == nil {
		return c.JSON(http.StatusFound, &user)
	}

	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}

func Create(c echo.Context) error {
	user := new(models.User)

	c.Bind(user)
	err := c.Validate(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	db, _ := c.Get("db").(pgsql.Storage)
	err = db.Save(user)

	if err == nil {
		return c.JSON(http.StatusCreated, user)
	}

	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
func Update(c echo.Context) error {
	user := new(models.User)

	db, _ := c.Get("db").(pgsql.Storage)

	c.Bind(user)
	err := c.Validate(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	oldLogin := c.Param("login")
	err = db.Change(oldLogin, user)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	} else {

		return c.JSON(http.StatusAccepted, user)

	}

	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
func Delete(c echo.Context) error {
	db, _ := c.Get("db").(pgsql.Storage)

	login := c.Param("login")
	err := db.Remove(login)

	if err == nil {
		return c.NoContent(http.StatusOK)
	}

	return echo.NewHTTPError(http.StatusNotFound, err.Error())
}
