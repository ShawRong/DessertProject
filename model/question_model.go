package model

type Question struct {
	Num  string `json:"num"`
	Rank string `json:"rank"`
	X    string `json:"x"`
	Y    string `json:"y"`
	Sig  string `json:"sig"`
	Res  string `json:"res"`
}
