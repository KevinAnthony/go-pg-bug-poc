package main

import (
	"github.com/go-pg/pg"
)

func setupSql() *pg.DB {
	return pg.Connect(&options)
}
