package main

import "github.com/go-pg/pg"

var options = pg.Options{
	Addr:     "localhost:5432",
	User:     "postgres",
	Password: "postgres",
	Database: "test",
}
