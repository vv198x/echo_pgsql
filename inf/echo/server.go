package echo

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
	"userSL/inf/pgsql"
	"userSL/pkg/cfg"
)

func Start(addr string) {
	e := echo.New()

	db := pgsql.GetPostgre()
	defer log.Println(db.CloseDB())

	//DB в контекст, один пулл и одно соединение на всё API
	e.Use(ContextDB(db))

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	if !cfg.Get().Debug {
		e.Use(middleware.Recover())
	}

	//go-playground validator
	e.Validator = &CustomValidator{Validator: validator.New()}

	//Настройки CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	//Хендлеры в отдельном файле headlers
	headlers(e)

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
