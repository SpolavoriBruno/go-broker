package entity

type Transaction struct {
	SellOrder *Order
	BuyOrder  *Order
	Price     float32
	Shares    int
	Total     float32
}

func NewTrasaction(sellOrder, buyOrder *Order, price float32, shares int) *Transaction {
	return &Transaction{
		SellOrder: sellOrder,
		BuyOrder:  buyOrder,
		Price:     price,
		Shares:    shares,
	}
}
