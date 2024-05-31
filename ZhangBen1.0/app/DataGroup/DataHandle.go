package DataGroup

import (
	"ZhangBen1.0/app/DataGroup/AddData"
	"ZhangBen1.0/app/DataGroup/DeleteData"
	"ZhangBen1.0/app/DataGroup/GetData"
	"github.com/gin-gonic/gin"
)

type DataHandle struct {
}

func (d *DataHandle) RegisterDataRoutes(r *gin.Engine) {
	DataGroup := r.Group("/Data")
	DataGroup.GET("/GetData", GetData.GetData)
}

type CURDHandle struct {
}

func (c *CURDHandle) RegisterCURDRoutes(r *gin.Engine) {
	ZDCURDGroup := r.Group("/CURD")
	ZDCURDGroup.POST("/AddZhangDan", AddData.AddZhangDan)
	ZDCURDGroup.POST("/delete-data", DeleteData.DeleteData)
}
