package client

type Payment int

const (
	Cash Payment = iota
	Card
	Paypal
)
