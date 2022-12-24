package migrations

import (
	"flag"
	"fmt"
	"github.com/go-pg/migrations"
	"github.com/go-pg/pg"
	"os"
	"userSL/pkg/cfg"
)

const usageText = `This program runs command on the db. Supported commands are:
  - init - creates version info table in the database
  - up - runs all available migrations.
  - up [target] - runs available migrations up to the target one.
  - down - reverts last migration.
  - reset - reverts all migrations.
  - version - prints current db version.
  - set_version [version] - sets db version without running migrations.
Usage:
  userSL <command> [args]
`

func Start() {
	flag.Usage = usage
	flag.Parse()
	arg := flag.Args()

	db := pg.Connect(&pg.Options{
		User:     cfg.Get().PGUser,
		Password: cfg.Get().PGPass,
		Addr:     cfg.Get().PGAddr,
		Database: cfg.Get().PGDB,
		PoolSize: 50,
	})
	defer db.Close()

	oldVersion, newVersion, err := migrations.Run(db, arg...)
	if err != nil {
		exitf(err.Error())
	}
	if newVersion != oldVersion {
		fmt.Printf("migrated from version %d to %d\n", oldVersion, newVersion)
	} else {
		fmt.Printf("version is %d\n", oldVersion)
	}
	if len(arg) > 0 && arg[0] == "init" {
		os.Exit(2)
	}
}

func usage() {
	fmt.Print(usageText)
	flag.PrintDefaults()
	os.Exit(2)
}

func errorf(s string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, s+"\n", args...)
}

func exitf(s string, args ...interface{}) {
	errorf(s, args...)
	os.Exit(1)
}
