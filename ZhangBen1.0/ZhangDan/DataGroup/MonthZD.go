package DataGroup

import (
	"ZhangBen1.0/DB"
	ZB "ZhangBen1.0/UserType"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MonthZD struct {
	YearMonth    string `json:"yearmonth"`
	MonthIncome  int    `json:"monthincome"`
	MonthExpense int    `json:"monthexpense"`
	Credit       int    `json:"credit"`
}

type MapZD struct {
	MonthIncome  int
	MonthExpense int
	Credit       int
}

func MonthData(c *gin.Context) {
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

	//以下求月账单
	var GetYearMonthData map[string]MapZD
	GetYearMonthData = make(map[string]MapZD)
	for _, r := range DBData {
		YearMonth := strconv.Itoa(r.Year) + "-" + strconv.Itoa(r.Month)
		if r.InEx == "expense" {
			t := GetYearMonthData[YearMonth]
			t.MonthExpense += r.Money
			t.Credit -= r.Money
			GetYearMonthData[YearMonth] = t
		}
		if r.InEx == "income" {
			t := GetYearMonthData[YearMonth]
			t.MonthIncome += r.Money
			t.Credit += r.Money
			GetYearMonthData[YearMonth] = t
		}
	}

	var MonthDT []MonthZD
	for k, v := range GetYearMonthData {
		var t MonthZD
		t.YearMonth = k
		t.MonthIncome = v.MonthIncome
		t.MonthExpense = v.MonthExpense
		t.Credit = v.Credit
		MonthDT = append(MonthDT, t)
	}

	c.JSON(http.StatusOK, MonthDT)
}
