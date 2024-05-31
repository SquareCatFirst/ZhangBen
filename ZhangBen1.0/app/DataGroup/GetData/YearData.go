package GetData

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
