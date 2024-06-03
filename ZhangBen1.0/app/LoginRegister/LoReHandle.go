package LoginRegister

import (
	"ZhangBen1.0/app/LoginRegister/Login"
	"ZhangBen1.0/app/LoginRegister/Register"
	"github.com/gin-gonic/gin"
)

type LoReHandle struct {
}

func (l *LoReHandle) RegisterLoReRoutes(r *gin.Engine) {

	LGRGGroup := r.Group("/LGRG")
	LGRGGroup.POST("/loginTry", Login.Login)
	LGRGGroup.POST("/reg", Register.Register)
	LGRGGroup.GET("/Captcha", Register.Captcha)
}
