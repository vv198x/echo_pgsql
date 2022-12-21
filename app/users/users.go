package users

import (
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
	if login == "" {
		users, err := db.LoadAll()
		if err != nil {
			log.Println("DB error ", err)
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, &users)
	}

	user, err := db.Load(login)
	if err == nil {
		return c.JSON(http.StatusOK, &user)
	}

	return echo.NewHTTPError(http.StatusNotFound, err.Error())
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
// @Failure	500
// @Router /{login} [delete]
func Delete(c echo.Context) error {
	db, _ := c.Get("db").(pgsql.Storage)

	login := c.Param("login")
	err := db.Remove(login)

	if err == nil {
		return c.NoContent(http.StatusOK)
	}

	return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
}
