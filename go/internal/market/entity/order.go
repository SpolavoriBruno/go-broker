package entity

type Order struct {
	Id            string
	Investor      *Investor
	Asset         *Asset
	Price         float32
	Shares        int
	Side          string
	Status        string
	OrderType     string
	PendingShares int // CLOSED | OPEN
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
