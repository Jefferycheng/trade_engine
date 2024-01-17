package domain

type SellOrderRepository interface {
	Empty() bool
	Add(o *Order)
	GetMinPriceOrder() (int, *Order)
	Remove(index int)
	GetAll() []*Order
}

// SellOrderMemoryRepository simulate database
type SellOrderMemoryRepository struct {
	orders []*Order
}

func NewMemorySellOrderRepository() *SellOrderMemoryRepository {
	return &SellOrderMemoryRepository{orders: make([]*Order, 0)}
}

func (r *SellOrderMemoryRepository) Empty() bool {
	return len(r.orders) == 0
}

func (r *SellOrderMemoryRepository) Add(order *Order) {
	r.orders = append(r.orders, order)
}

func (r *SellOrderMemoryRepository) Remove(index int) {
	r.orders = append(r.orders[:index], r.orders[index+1:]...)
}

func (r *SellOrderMemoryRepository) GetMinPriceOrder() (int, *Order) {
	if len(r.orders) == 0 {
		return -1, nil
	}

	minOrder := r.orders[0]

	var minIndex = 0

	for i, o := range r.orders {
		if minOrder.Price > o.Price {
			minOrder = o
			minIndex = i
		}
	}

	return minIndex, minOrder
}

func (r *SellOrderMemoryRepository) GetAll() []*Order {
	return r.orders
}
