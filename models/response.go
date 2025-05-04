package models

type ListResponse struct {
	Items []interface{} `json:"items"`
	Count int                   `json:"count"`
	Page  int                   `json:"page"`
}