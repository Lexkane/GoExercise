package orders

import (
	"algorithm/algorithm/client"
	"algorithm/algorithm/ticket"
)

type Orders struct {
	orders map[client.Client][]ticket.Ticket
}
