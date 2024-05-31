package LoginRegister

import "github.com/gin-gonic/gin"

type LoReHandle struct {
}

func (l *LoReHandle) RegisterLoReRoutes(r *gin.Engine) {
	LGRGGroup := r.Group("/LGRG")
	LGRGGroup.POST("/loginTry", Login)
	LGRGGroup.POST("/reg", Register)
}
