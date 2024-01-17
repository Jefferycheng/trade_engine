package domain

type Order struct {
	BuyOrSell bool // true is buy, false is sell
	Quantity  int
	Price     int // price less than or equal to 0, define is market price
}

func (o Order) IsMarketPrice() bool {
	return o.Price <= 0
}
