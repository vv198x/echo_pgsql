package main

import (
	"fmt"
	_ "github.com/swaggo/files"
	"golang.org/x/crypto/bcrypt"
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
//	@tag.description Read only access. For all authorized but not blocked
//	@tag.name	auth
//	@tag.description For authorization

// @securityDefinitions.basic  BasicAuth
func main() {
	config.Load()
	logger.Start()
	pgsql.ReplaceTable(*config.SQLScript)
	//echo.Start(*config.Address)
	buf, _ := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
	fmt.Println(string(buf))
	buf, _ = bcrypt.GenerateFromPassword([]byte("user"), bcrypt.DefaultCost)
	fmt.Println(string(buf))
	buf, _ = bcrypt.GenerateFromPassword([]byte("lock"), bcrypt.DefaultCost)
	fmt.Println(string(buf))
}
