package route

import ()

type Route struct {
	ID           int
	name         string
	class        TypeOfRide
	startStation string
	endStation   string
	path         []string
	booked       map[string][][]bool
}

type Train struct {
	places [][]bool
}

//var AllRoutes map[string]Route

func init() {

}
