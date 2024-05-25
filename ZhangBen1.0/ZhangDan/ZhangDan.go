package ZhangDan

import (
	"ZhangBen1.0/DB"
	ZB "ZhangBen1.0/UserType"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Res struct {
	Book  string `json:"book"`
	Typ   string `json:"typ"`
	InEx  string `json:"inex"`
	Money string `json:"money"`
	Note  string `json:"note"`
}
type Book struct {
	Option string `json:"book"`
}

func SetZhangBen(c *gin.Context) {

}

func AddZhangDan(c *gin.Context) {
	var ZD Res
	if err := c.ShouldBindJSON(&ZD); err != nil {
		fmt.Println("有问题")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := c.Cookie("uid")
	if err != nil {
		fmt.Println("从cookie获取uid失败")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	mon, err := strconv.Atoi(ZD.Money)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := strconv.Atoi(uid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	b, err := strconv.Atoi(ZD.Book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentTime := time.Now()
	year, month, day := currentTime.Date()
	Month := int(month)

	ZhangDan := &ZB.ZhangDan{
		Uid:   u,
		Book:  b,
		Typ:   ZD.Typ,
		InEx:  ZD.InEx,
		Money: mon,
		Note:  ZD.Note,
		Year:  year,
		Month: Month,
		Day:   day,
	}
	err = DB.Db.Insert(ZhangDan)
	if err != nil {
		fmt.Println("向数据库中插入数据失败")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Data received"})
}
