package entity

import (
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	SellOrder *Order
	BuyOrder  *Order
	Id        string
	Price     float32
	Shares    int
	Total     float32
	OpenTime  time.Time
	CloseTime time.Time
}

func NewTrasaction(buyOrder, sellOrder *Order) *Transaction {
	asset := sellOrder.Asset
	price := sellOrder.Price
	shares := int(math.Min(float64(sellOrder.PendingShares), float64(buyOrder.PendingShares)))

	fmt.Println("Update Orders")
	sellOrder.UpdatePendingShares(-shares)
	buyOrder.UpdatePendingShares(-shares)

	sellOrder.CloseOrder()
	buyOrder.CloseOrder()

	fmt.Println("Update Asset")
	asset.Price = price

	fmt.Println("Update Investor")
	sellOrder.Investor.UpdateAssetPosition(asset.Id, -shares)
	buyOrder.Investor.UpdateAssetPosition(asset.Id, shares)

	fmt.Println("New transaction created")

	return &Transaction{
		SellOrder: sellOrder,
		BuyOrder:  buyOrder,
		Price:     price,
		Shares:    shares,
		Id:        uuid.New().String(),
		OpenTime:  time.Now(),
	}
}

/** Whats hapend with this header? */
func (t *Transaction) CalculateTotalPrice(shares int, price float32) float32 {
	t.Total = t.Price * float32(t.Shares)
	return t.Total
}
