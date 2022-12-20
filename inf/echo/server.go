package echo

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"userSL/inf/pgsql"
)

func Start(addr string) {
	e := echo.New()

	db := pgsql.GetPostgre()
	defer db.CloseDB()

	e.Use(ContextDB(db))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
		//Output: logger.LogFile,
	}))
	e.Use(middleware.Recover())
	e.Validator = &CustomValidator{Validator: validator.New()}

	routes(e)

	e.Logger.Fatal(e.Start(addr))
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func ContextDB(db pgsql.Storage) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
