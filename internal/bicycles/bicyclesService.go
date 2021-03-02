package goRest

import . "github.com/lucas-gio/goRest/internal/interfaces"

type BicyclesService struct {
	Datasource IDatasourceService `di.inject:"datasourceService"`
}

/*listBicycles returns an slice of bicycles that match the criteria to find. */
/*func (b *BicyclesService) listBicycles() (bicyclesList []Bicycle, err error){

}*/
