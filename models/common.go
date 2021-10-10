package models

type Page struct {
	Offset int         `json:"offset"`
	Limit  int         `json:"limit"`
	Data   interface{} `json:"data"`
	Total  int64       `json:"total"`
}
