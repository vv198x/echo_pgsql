package main

import (
	"userSL/inf/echo"
	"userSL/pkg/logger"
)

func main() {
	logger.Start()
	//pgsql.ReplaceTable("./table.sql")
	echo.Start(":8000")
}
