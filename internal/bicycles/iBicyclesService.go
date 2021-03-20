package goRest

type IBicyclesService interface {
	ListBicycles(string, string) (bicyclesList *[]Bicycle)
}
