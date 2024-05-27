package DataGroup

import (
	"ZhangBen1.0/DB"
	ZB "ZhangBen1.0/UserType"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HistoryZD struct {
	Dataid       int    `json:"dataid"`
	Typ          string `json:"typ"`
	Money        int    `json:"money"`
	Notes        string `json:"notes"`
	YearMonthDay string `json:"yearmonthday"`
}

func HistoryData(c *gin.Context) {
	uid, err := c.Cookie("uid")
	if err != nil {
		fmt.Println("从cookie获取uid失败")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var DBData []ZB.ZhangDan
	var HistoryDT []HistoryZD
	err = DB.Db.Model(&DBData).Where("uid = ?", uid).Select()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//	fmt.Println(DBData)
	for _, r := range DBData {
		var ZD HistoryZD
		switch r.Typ {
		case "food":
			ZD.Typ = "餐饮"
		case "daily":
			ZD.Typ = "日用"
		case "transportation":
			ZD.Typ = "交通"
		case "sport":
			ZD.Typ = "运动"
		case "study":
			ZD.Typ = "学习"
		case "donation":
			ZD.Typ = "捐赠"
		case "envelope":
			ZD.Typ = "红包"
		case "game":
			ZD.Typ = "游戏"
		case "other":
			ZD.Typ = "其他"
		default:
			ZD.Typ = r.Typ
		}
		switch r.InEx {
		case "expense":
			ZD.Money = -r.Money
		case "income":
			ZD.Money = r.Money
		default:
			ZD.Money = -1
		}
		ZD.Dataid = r.Dataid
		ZD.Notes = r.Note
		ZD.YearMonthDay = strconv.Itoa(r.Year) + "-" + strconv.Itoa(r.Month) + "-" + strconv.Itoa(r.Day)
		HistoryDT = append(HistoryDT, ZD)
	}
	//fmt.Println(HistoryDT)
	c.JSON(http.StatusOK, HistoryDT)
}

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
