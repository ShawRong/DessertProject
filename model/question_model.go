package model

type Question struct {
	Num     string `json:"num"`
	Rank    string `json:"rank"`
	Content string `json:"content"`
	Res     string `json:"res"`
}
