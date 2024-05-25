package ZBtypes

type ZhangDan struct {
	Dataid int    `json:"dataid"`
	Uid    int    `json:"uid"`
	Book   int    `json:"book"`
	Typ    string `json:"typ"`
	InEx   string `json:"inex"`
	Money  int    `json:"money"`
	Note   string `json:"note"`
	Year   int    `json:"year"`
	Month  int    `json:"month"`
	Day    int    `json:"day"`
}

type ZhangBen struct {
	Uid int        `json:"uid"`
	Zd  []ZhangDan `json:"zd"`
}
