package entity

import "fmt"

type Order struct {
	Id            string
	Investor      *Investor
	Asset         *Asset
	Price         float32
	Shares        int
	Side          string
	Status        string // CLOSED | OPEN
	OrderType     string
	PendingShares int
}

func NewOrder(id string, investor *Investor, asset *Asset, shares int, price float32, side string) *Order {
	return &Order{
		Id:            id,
		Investor:      investor,
		Asset:         asset,
		Price:         price,
		Shares:        shares,
		OrderType:     side,
		PendingShares: shares,
		Status:        "OPEN",
	}
}

func (o *Order) CloseOrder() {
	if o.PendingShares == 0 {
		o.Status = "CLOSED"
		fmt.Println("Order closed")
	}
}

func (o *Order) UpdatePendingShares(shares int) {
	o.PendingShares += shares
}
