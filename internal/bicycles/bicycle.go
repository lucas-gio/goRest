package goRest

type Bicycle struct {
	Id          string
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Origin      string `json:"origin"`
	UnitPrice   string `json:"unitPrice"`
	Image       string `json:"image"`
}
