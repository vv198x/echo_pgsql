package main

import (
	"usersSL/inf/echo"
	"usersSL/pkg/logger"
)

func main() {
	logger.Start()
	echo.Start(":8000")
}
