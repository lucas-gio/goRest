package models

type Bicycle struct {
	ID          string
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
