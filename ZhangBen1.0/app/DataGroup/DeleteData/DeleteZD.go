package DeleteData

import (
	"ZhangBen1.0/DB"
	ZB "ZhangBen1.0/UserType"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Request struct {
	DataId int `json:"dataid"`
}

func DeleteData(c *gin.Context) {
	var id Request
	if err := c.ShouldBindJSON(&id); err != nil {
		fmt.Println("绑定json失败")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := DB.Db.Model(&ZB.ZhangDan{}).Where("dataid = ?", id.DataId).Delete(); err != nil {
		fmt.Println("删除历史数据失败")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
