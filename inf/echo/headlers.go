package echo

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func headlers(e *echo.Echo) {
	root := e.Group("/api/users/v1")
	login := e.Group("/api/users/v1/:login")

	root.POST("/", Create, middleware.BasicAuth(forAdmin))
	login.PUT("", Update, middleware.BasicAuth(forAdmin))
	login.DELETE("", Delete, checkLastAdmin, middleware.BasicAuth(forAdmin))
	// Для пакета swag разделил Read. И не получилось в одну совместить
	root.GET("/", ReadAll, middleware.BasicAuth(forAll))
	login.GET("", Read, middleware.BasicAuth(forAll))

}
