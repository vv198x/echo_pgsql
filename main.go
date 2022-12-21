package main

import (
	_ "github.com/swaggo/files"
	"userSL/inf/echo"
	"userSL/inf/pgsql"
	"userSL/pkg/logger"
)

// @title           Simple API
// @version         0.0.1
// @description     Api server

// @host      80.92.206.187:8000
// @BasePath  /api/users/v1
// @schemes http
//	@tag.name	admins
//	@tag.description Admin access
//	@tag.name	read
//	@tag.description Read only access. For all but blocked

// @securityDefinitions.basic  BasicAuth
func main() {
	logger.Start()
	pgsql.ReplaceTable("./table.sql")
	echo.Start(":8000")
}
