package ridemanager

import (
	"algorithm/algorithm/route"
	"algorithm/algorithm/station"
)

var AllRides map[string]route.Route
var bookedplaces map[station.Station]route.Train

// Map of key valuse pairs "Lviv-Kyiv" = route type from Lviv to Kyiv
