package main

import (
	"fmt"
	"github.com/go-pg/pg"
)

func setupSql() *pg.DB {
	host := ""
	port := 0
	user := ""
	pass := ""
	database := ""
	options := pg.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		User:     user,
		Password: pass,
		Database: database,
	}

	return pg.Connect(&options)
}
