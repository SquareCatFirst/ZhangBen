package DB

import "github.com/go-pg/pg"

func DBinit() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "192.168.137.139:5432",
		User:     "postgres",
		Password: "114514",
		Database: "ZB",
	})
	return db
}

var Db *pg.DB
