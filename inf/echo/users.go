package echo

import (
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"strings"
	"userSL/inf/pgsql"
	"userSL/models"
)

// Read godoc
// @Tags read
// @Summary Retrieves user based on given Login
// @Produce json
// @Param login path string true "User login"
// @Success 200 {object} models.User
// @Failure	404 {object} models.JSONResult{message=string} "Not found"
// @Failure	500
// @Router /{login} [get]
func Read(c echo.Context) error {
	db, _ := c.Get("db").(pgsql.Storage)

	login := c.Param("login")

	user, err := db.Load(login)
	//Не выводить пароль. Так занулил, если тип сменится.
	user.Password = models.User{}.Password
	if err == nil && len(user.Login) != 0 {
		return c.JSON(http.StatusOK, &user)
	} else {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	log.Println("DB error", err)

	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

// Read godoc
// @Tags read
// @Summary Retrieves users
// @Produce json
// @Success 200 {object} models.User
// @Failure	500
// @Router / [get]
func ReadAll(c echo.Context) error {
	fmt.Println(c.Param("login"))
	db, _ := c.Get("db").(pgsql.Storage)
	users, err := db.LoadAll()
	if err == nil {
		for i := range users {
			users[i].Password = models.User{}.Password
		}
		return c.JSON(http.StatusOK, &users)
	}

	log.Println("DB error ", err)
	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

// Create godoc
// @Summary Create new user
// @Tags admins
// @Produce json
// @Param message body models.User true  "New user"
// @Success 201 {object} models.User
// @Failure	400 {object} models.JSONResult{message=string} "Validation error"
// @Failure	409 {object} models.JSONResult{message=string} "User with this login exists"
// @Failure	500
// @Router / [post]
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

	if strings.Contains(err.Error(), "duplicate") {
		return echo.NewHTTPError(http.StatusConflict, err.Error())
	}

	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}

// Update godoc
// @Summary Update user on given Login
// @Tags admins
// @Produce json
// @Param login path string true "User login"
// @Param message body models.User true  "Update user"
// @Success 202 {object} models.User
// @Failure	404 {object} models.JSONResult{message=string} "Not found"
// @Failure	400 {object} models.JSONResult{message=string} "Validation error"
// @Failure	500
// @Router /{login} [put]
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

// Delete godoc
// @Tags admins
// @Summary Delete user on given Login
// @Produce json
// @Param login path string true "User login"
// @Success 200
// @Failure	400 {object} models.JSONResult{message=string} "Attempt to remove the last admin"
// @Failure	404 {object} models.JSONResult{message=string} "Not found"
// @Failure	500
// @Router /{login} [delete]
func Delete(c echo.Context) error {
	db, _ := c.Get("db").(pgsql.Storage)
	user := *(c.Get("user").(*models.User))

	err := db.Remove(user.Login, user.Rule)

	if err == nil {
		return c.NoContent(http.StatusOK)
	}

	return echo.NewHTTPError(http.StatusOK, err.Error())
}
