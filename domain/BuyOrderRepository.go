package domain

type BuyOrderRepository interface {
	Empty() bool
	Add(o *Order)
	GetMaxPriceOrder() (int, *Order)
	Remove(index int)
	GetAll() []*Order
}

// BuyOrderMemoryRepository simulate database
type BuyOrderMemoryRepository struct {
	orders []*Order
}

func NewMemoryBuyOrderRepository() *BuyOrderMemoryRepository {
	return &BuyOrderMemoryRepository{orders: make([]*Order, 0)}
}

func (r *BuyOrderMemoryRepository) Empty() bool {
	return len(r.orders) == 0
}

func (r *BuyOrderMemoryRepository) Add(order *Order) {
	r.orders = append(r.orders, order)
}

func (r *BuyOrderMemoryRepository) Remove(index int) {
	r.orders = append(r.orders[:index], r.orders[index+1:]...)
}

func (r *BuyOrderMemoryRepository) GetMaxPriceOrder() (int, *Order) {
	if len(r.orders) == 0 {
		return -1, nil
	}

	maxOrder := r.orders[0]

	var maxIndex = 0

	for i, o := range r.orders {
		if maxOrder.Price < o.Price {
			maxOrder = o
			maxIndex = i
		}
	}

	return maxIndex, maxOrder
}

func (r *BuyOrderMemoryRepository) GetAll() []*Order {
	return r.orders
}
