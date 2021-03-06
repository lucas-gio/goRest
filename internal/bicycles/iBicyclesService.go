package goRest

type IBicyclesService interface {
	ListBicycles() (bicyclesList *[]Bicycle)
}
