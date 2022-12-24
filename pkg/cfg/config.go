package cfg

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	JWTkey    string `env:"JWTkey" env-default:"SECRET" env-description:"Secret for JWT"`
	Address   string `env:"addr" env-default:":8000" env-description:"Address and port for echo"`
	SQLScript string `env:"SQLScript" env-default:"./Docker/table.sql" env-description:"Script for first load"`
	PGUser    string `env:"PGUser" env-default:"pgsql" env-description:"User name for PostgreSQL"`
	PGPass    string `env:"PGPass" env-default:"PASS" env-description:"Password for PostgreSQL"`
	PGAddr    string `env:"PGAddr" env-default:"postgres:5432" env-description:"Password for PostgreSQL"`
	PGDB      string `env:"PGDB" env-default:"pgdb" env-description:"Name database for PostgreSQL"`
	Debug     bool   `env:"Debug" env-default:0 env-description:"Debug start mode"`
}

var c Config

func init() {
	err := cleanenv.ReadEnv(&c)
	if err != nil {
		log.Fatalln("Can't create config")
	}
}

func Get() *Config {
	return &c
}
