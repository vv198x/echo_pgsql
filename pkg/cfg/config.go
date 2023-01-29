package cfg

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	JWTkey    string `env:"JWTkey" env-default:"SECRET" env-description:"Secret for JWT"`
	Address   string `env:"addr" env-default:":8000" env-description:"Address and port for echo"`
	PGUser    string `env:"PGUser" env-default:"postgres" env-description:"User name for PostgreSQL"`
	PGPass    string `env:"PGPass" env-default:"password" env-description:"Password for PostgreSQL"`
	PGAddr    string `env:"PGAddr" env-default:"192.168.122.161:5432" env-description:"Password for PostgreSQL"`
	PGDB      string `env:"PGDB" env-default:"database" env-description:"Name database for PostgreSQL"`
	Migration string `env:"Migration" env-default:"" env-description:"To initiate a migration. init"`
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
