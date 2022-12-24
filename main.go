package main

import (
	_ "github.com/swaggo/files"
	"userSL/inf/echo"
	"userSL/migrations"
	"userSL/pkg/cfg"
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
//	@tag.description Read only access. For all authorized but not blocked
//	@tag.name	auth
//	@tag.description For authorization

// @securityDefinitions.basic  BasicAuth
func main() {
	logger.Start()
	migrations.Start()
	//pgsql.ReplaceTable(cfg.Get().SQLScript)
	echo.Start(cfg.Get().Address)
}
