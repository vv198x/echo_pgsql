package main

import (
	_ "github.com/swaggo/files"
	"userSL/inf/echo"
	"userSL/migrations"
	"userSL/pkg/cfg"
	"userSL/pkg/logger"
)

// @title           Simple API
// @version         0.0.9
// @description     Api server

// @host      localhost:8000
// @BasePath  /api/users/v1
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
//	@in	query
//	@name token
//	@tag.name	admins
//	@tag.description Admin access
//	@tag.name	read
//	@tag.description Read only access. For all authorized but not blocked
//	@tag.name	auth
//	@tag.description For authorization

func main() {
	logger.Start()
	migrations.Start()
	echo.Start(cfg.Get().Address)
}
