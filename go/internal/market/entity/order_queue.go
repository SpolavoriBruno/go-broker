package entity

type OrderQueue struct {
	Orders []*Order
}

func NewOrderQueue() *OrderQueue {
	return &OrderQueue{}
}

// Len implements heap.Interface.
func (oq *OrderQueue) Len() int {
	return len(oq.Orders)
}

// Less implements heap.Interface.
func (oq *OrderQueue) Less(i int, j int) bool {
	return oq.Orders[i].Price < oq.Orders[j].Price
}

// Pop implements heap.Interface.
func (oq *OrderQueue) Pop() any {
	old := oq.Orders
	n := len(old)
	item := old[n-1]
	oq.Orders = old[0 : n-1]
	return item
}

// Push implements heap.Interface.
func (oq *OrderQueue) Push(x any) {
	oq.Orders = append(oq.Orders, x.(*Order))
}

// Swap implements heap.Interface.
func (oq *OrderQueue) Swap(i int, j int) {
	oq.Orders[i], oq.Orders[j] = oq.Orders[j], oq.Orders[i]
}
