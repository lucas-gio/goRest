package goRest

type iPartsService interface {
	ListParts(string, string) (partsList *[]Part)
}
