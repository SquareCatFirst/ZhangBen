package DataGroup

import (
	"ZhangBen1.0/DB"
	ZB "ZhangBen1.0/UserType"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type YearZD struct {
	Year        string `json:"years"`
	YearIncome  int    `json:"yearincome"`
	YearExpense int    `json:"yearexpense"`
	Credit      int    `json:"credit"`
}

type MapZDY struct {
	YearIncome  int
	YearExpense int
	Credit      int
}

func YearData(c *gin.Context) {
	uid, err := c.Cookie("uid")
	if err != nil {
		fmt.Println("从cookie获取uid失败")
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var DBData []ZB.ZhangDan
	err = DB.Db.Model(&DBData).Where("uid = ?", uid).Select()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//以下求年账单
	var GetYearData map[string]MapZDY
	GetYearData = make(map[string]MapZDY)
	for _, r := range DBData {
		Year := strconv.Itoa(r.Year)
		if r.InEx == "expense" {
			t := GetYearData[Year]
			t.YearExpense += r.Money
			t.Credit -= r.Money
			GetYearData[Year] = t
		}
		if r.InEx == "income" {
			t := GetYearData[Year]
			t.YearIncome += r.Money
			t.Credit += r.Money
			GetYearData[Year] = t
		}
	}

	var YearDT []YearZD
	for k, v := range GetYearData {
		var t YearZD
		t.Year = k
		t.YearIncome = v.YearIncome
		t.YearExpense = v.YearExpense
		t.Credit = v.Credit
		YearDT = append(YearDT, t)
	}

	c.JSON(http.StatusOK, YearDT)
}
