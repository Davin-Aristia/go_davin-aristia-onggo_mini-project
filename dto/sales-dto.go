package dto

import "time"

type SalesRequest struct {
	UserId uint `json:"user" form:"user"`
	Invoice string `json:"invoice" form:"invoice"`
	Date time.Time `json:"date" form:"date"`
	Total float64 `json:"total" form:"total"`
}