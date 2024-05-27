package DataGroup

import (
	"github.com/gin-gonic/gin"
)

type DataHandle struct {
}

func (d *DataHandle) RegisterDataRoutes(r *gin.Engine) {
	DataGroup := r.Group("/Data")
	DataGroup.GET("/HistoryData", HistoryData)
	DataGroup.GET("/MonthData", MonthData)
	DataGroup.GET("/YearData", YearData)
}

type CURDHandle struct {
}

func (c *CURDHandle) RegisterCURDRoutes(r *gin.Engine) {
	ZDCURDGroup := r.Group("/CURD")
	ZDCURDGroup.POST("/AddZhangDan", AddZhangDan)
	ZDCURDGroup.POST("/delete-data", DeleteData)
}
