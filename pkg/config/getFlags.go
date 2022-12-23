package config

import (
	"flag"
)

// Не в переменных окружения. И не в файле
// Потому что - сразу дефолтные значения можно указать и быстрее.
var (
	Address    = flag.String("addr", ":8000", "Address and port for echo")
	SQLScript  = flag.String("SQLScript", "./Docker/table.sql", "Script for first load")
	PGUser     = flag.String("PGUser", "pgsql", "User name for PostgreSQL")
	PGPass     = flag.String("PGPass", "PASS", "Password for PostgreSQL")
	PGAddr     = flag.String("PGAddr", "postgres:5432", "Password for PostgreSQL")
	PGDB       = flag.String("PGDB", "pgdb", "Name database for PostgreSQL")
	PGPoolSize = flag.Int("PGPoolSize", 50, "Pool size")
	Debug      = flag.Bool("debug", false, "Debug start mode")
)

// Для го тестов. Парсинг нужно вынести из init
func Load() {
	flag.Parse()
}
