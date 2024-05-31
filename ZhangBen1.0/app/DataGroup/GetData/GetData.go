package GetData

import (
	"ZhangBen1.0/DB"
	ZB "ZhangBen1.0/UserType"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type DataZD struct {
	HistoryZD []HistoryZD `json:"historyzd"`
	MonthZD   []MonthZD   `json:"monthzd"`
	YearZD    []YearZD    `json:"yearzd"`
}

func GetData(c *gin.Context) {
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

	//以下求历史账单
	var HistoryDT []HistoryZD
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

	Data := DataZD{
		HistoryZD: HistoryDT,
		MonthZD:   MonthDT,
		YearZD:    YearDT,
	}
	c.JSON(http.StatusOK, Data)
}
