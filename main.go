package main

import (
	"userSL/inf/echo"
	"userSL/pkg/logger"
)

func main() {
	logger.Start()
	echo.Start(":8000")
}
