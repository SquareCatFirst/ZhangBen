package GetData

type HistoryZD struct {
	Dataid       int    `json:"dataid"`
	Typ          string `json:"typ"`
	Money        int    `json:"money"`
	Notes        string `json:"notes"`
	YearMonthDay string `json:"yearmonthday"`
}
