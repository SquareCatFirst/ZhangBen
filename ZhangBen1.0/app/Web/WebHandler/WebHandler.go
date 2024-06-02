package WebHandler

import (
	DG "ZhangBen1.0/app/DataGroup"
	LoRe "ZhangBen1.0/app/LoginRegister"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type WebHandler struct {
}

func (*WebHandler) Routes(r *gin.Engine) {
	r.Static("/test", "./templates")
	r.LoadHTMLGlob("templates/*")

	r.GET("", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/test/login.html")
	})
	r.GET("/login", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/test/login.html")
	})

	d := DG.DataHandle{}
	d.RegisterDataRoutes(r)

	c := DG.CURDHandle{}
	c.RegisterCURDRoutes(r)

	l := LoRe.LoReHandle{}
	l.RegisterLoReRoutes(r)

	//	router.POST("/ZB-set-cookie")
	go func() {
		err := r.Run("10.0.1.104:1145")
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func (*WebHandler) ServerExit() {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting down server...")
	os.Exit(0)
}
