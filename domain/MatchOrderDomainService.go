package domain

type MatchOrderDomainService interface {
	MatchBuyOrder(buyOrder *Order) (bool, int)
	MatchSellOrder(sellOrder *Order) (bool, int)
}

type MatchOrderDS struct {
	SellOrderRepository SellOrderRepository
	BuyOrderRepository  BuyOrderRepository
}

func (mo MatchOrderDS) MatchBuyOrder(buyOrder *Order) (bool, int) {
	if mo.SellOrderRepository.Empty() {
		mo.BuyOrderRepository.Add(buyOrder)
		return false, 0
	}

	if buyOrder.IsMarketPrice() {
		// match the lowest price and lower timestamp
		minPriceOrderIndex, minOrder := mo.SellOrderRepository.GetMinPriceOrder()
		mo.SellOrderRepository.Remove(minPriceOrderIndex)
		return true, minOrder.Price
	}

	for index, sellOrder := range mo.SellOrderRepository.GetAll() {
		// match the equals price or lower price
		if sellOrder.Price <= buyOrder.Price {
			mo.SellOrderRepository.Remove(index)
			return true, buyOrder.Price
		}
	}

	// If not filled, queue into buy order list
	mo.BuyOrderRepository.Add(buyOrder)
	return false, 0
}

func (mo MatchOrderDS) MatchSellOrder(sellOrder *Order) (bool, int) {

	if mo.BuyOrderRepository.Empty() {
		mo.SellOrderRepository.Add(sellOrder)
		return false, 0
	}

	if sellOrder.IsMarketPrice() {
		// match the highest price and lower timestamp
		maxPriceOrderIndex, maxOrder := mo.BuyOrderRepository.GetMaxPriceOrder()
		mo.BuyOrderRepository.Remove(maxPriceOrderIndex)
		return true, maxOrder.Price
	}

	for index, buyOrder := range mo.BuyOrderRepository.GetAll() {
		// match the equals price or higher price
		if buyOrder.Price >= sellOrder.Price {
			mo.BuyOrderRepository.Remove(index)
			return true, sellOrder.Price
		}
	}

	// If not filled, queue into sell order list
	mo.SellOrderRepository.Add(sellOrder)
	return false, 0
}
