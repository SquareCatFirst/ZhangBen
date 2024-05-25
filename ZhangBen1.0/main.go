package main

import (
	"ZhangBen1.0/DB"
	LG "ZhangBen1.0/Login"
	ZD "ZhangBen1.0/ZhangDan"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func ShowLOGIN(c *gin.Context) {
	fmt.Println(c.ClientIP())
	c.HTML(200, "login.html", gin.H{})
}
func GoToLOGIN(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "login")
}
func ShowReg(c *gin.Context) {
	c.HTML(200, "register.html", gin.H{})
}
func ShowINDEX(c *gin.Context) {
	//c.HTML(200, "index.html", gin.H{})
}

func main() {
	DB.Db = DB.DBinit()
	defer DB.Db.Close()
	router := gin.Default()

	router.Static("/test", "./templates")
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/test/login.html")
	})
	router.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/test/login.html")
	})

	router.GET("/HistoryData", ZD.HistoryData)
	router.GET("/MonthData", ZD.MonthData)
	router.GET("/YearData", ZD.YearData)

	router.POST("/index/AddZhangDan", ZD.AddZhangDan)
	router.POST("/delete-data", ZD.DeleteData)

	router.POST("/loginTry", LG.Login)

	router.POST("/ZB-set-cookie")

	err := router.Run("10.0.1.104:1145")
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")
	os.Exit(0)

}
