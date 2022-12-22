package main

import (
	_ "github.com/swaggo/files"
	"userSL/inf/echo"
	"userSL/inf/pgsql"
	"userSL/pkg/config"
	"userSL/pkg/logger"
)

// @title           Simple API
// @version         0.0.1
// @description     Api server

// @host      localhost:8000
// @BasePath  /api/users/v1
// @schemes http
//	@tag.name	admins
//	@tag.description Admin access
//	@tag.name	read
//	@tag.description Read only access. For all but blocked

// @securityDefinitions.basic  BasicAuth
func main() {
	logger.Start()
	pgsql.ReplaceTable(*config.SQLScript)
	echo.Start(*config.Address)
}
