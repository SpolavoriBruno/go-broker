package entity

import (
	"sync"

	"container/heap"
)

type Book struct {
	Id              string
	OrderInputChan  chan *Order
	OrderOutputChan chan *Order
	Transactions    []Transaction
	Wg              *sync.WaitGroup
}

func NewBook(inChan, outChan chan *Order, wg *sync.WaitGroup) *Book {
	return &Book{
		OrderInputChan:  inChan,
		OrderOutputChan: outChan,
		Wg:              wg,
	}
}

func (b *Book) Trade() {
	buyOrderQueue := make(map[string]*OrderQueue)
	sellOrderQueue := make(map[string]*OrderQueue)

	for order := range b.OrderInputChan {
		asset := order.Asset.Id

		if buyOrderQueue[asset] == nil {
			buyOrderQueue[asset] = NewOrderQueue()
			heap.Init(buyOrderQueue[asset])
		}
		if sellOrderQueue[asset] == nil {
			sellOrderQueue[asset] = NewOrderQueue()
			heap.Init(sellOrderQueue[asset])
		}

		if order.OrderType == "BUY" {
			buyOrderQueue[asset].Push(order)
			if sellOrderQueue[asset].Len() > 0 {
				pair := sellOrderQueue[asset].Pop().(*Order)

				b.AddTransaction(order, pair, b.Wg)
			}
		} else if order.OrderType == "SELL" {
			sellOrderQueue[asset].Push(order)
			if buyOrderQueue[asset].Len() > 0 {
				pair := buyOrderQueue[asset].Pop().(*Order)

				b.AddTransaction(pair, order, b.Wg)
			}
		}
	}
}

func (b *Book) AddTransaction(buyOrder, sellOrder *Order, wg *sync.WaitGroup) {
	defer wg.Done()

	transaction := NewTrasaction(buyOrder, sellOrder)
	b.Transactions = append(b.Transactions, *transaction)
}
