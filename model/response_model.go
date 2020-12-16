package model

type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}
