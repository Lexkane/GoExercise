package ticket

import (
	"algorithm/algorithm/route"
)

type TicketBuilder struct {
	tickets Ticket
}

func generatePrice(carriage route.TypeOfCarriage, routeDistance int) float32 {
	var totalPrice float32 = 0

	if carriage == route.COMPARTMENT {
		totalPrice = float32(routeDistance / 3.0)
	}
	if carriage == route.BUSINESS {
		totalPrice = float32(routeDistance / 5.0)

	}

	if carriage == route.ECONOM {
		totalPrice = float32(routeDistance / 10.0)
	}

	return totalPrice

}
