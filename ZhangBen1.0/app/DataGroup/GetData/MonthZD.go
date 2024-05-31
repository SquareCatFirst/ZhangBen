package GetData

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
