package main

import (
	"ZhangBen1.0/DB"
	web "ZhangBen1.0/app/Web/WebHandler"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	DB.Db = DB.DBinit()
	defer DB.Db.Close()
	router := gin.Default()
	var w web.WebHandler
	w.Routes(router)
	w.ServerExit()
}
