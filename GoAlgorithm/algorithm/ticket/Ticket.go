package ticket

import (
	"algorithm/algorithm/client"
	"time"
)

type Ticket struct {
	name           string
	journeyName    string
	date           time.Time
	carriageNumber int
	place          int
	price          int
}

func newTicket(name string, journeyName string, date time.Time, carriageNumber int, place int, price int) *Ticket {
	return &Ticket{name: name, journeyName: journeyName, date: date, carriageNumber: carriageNumber, place: place, price: price}
}

var Tickets []Ticket

var TickesForUsers map[client.Client][]*Ticket
