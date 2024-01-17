package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"math/rand"
	"myapp/domain"
	"net/http"
	"strconv"
)

var channel chan domain.Order

var sellOrderRepo = domain.NewMemorySellOrderRepository()
var buyOrderRepo = domain.NewMemoryBuyOrderRepository()

// Create an instance of MatchOrderDS
var matchOrderService = &domain.MatchOrderDS{
	SellOrderRepository: sellOrderRepo,
	BuyOrderRepository:  buyOrderRepo,
}

func main() {
	channel = make(chan domain.Order)

	// In distributed system, replace with message queue
	go ReceiveChannel()

	e := echo.New()

	e.GET("/submit_order", func(c echo.Context) error {
		buyOrSell := c.QueryParam("b")
		price := c.QueryParam("p")
		count := c.QueryParam("c")

		b, _ := strconv.ParseBool(buyOrSell)
		p, _ := strconv.Atoi(price)
		cou, _ := strconv.Atoi(count)
		go sendOrderToChannel(b, p, cou)
		return c.String(http.StatusOK, "Order submitted")
	})

	e.GET("/concurrent_test", func(c echo.Context) error {
		for {
			// random buy or sell
			bs := rand.Intn(2)
			// random price
			r := rand.Intn(5)
			go sendOrderToChannel(bs%2 == 0, r, 1)
		}
	})

	e.Logger.Fatal(e.Start(":1323"))
}

func ReceiveChannel() {
	for {
		select {
		case order, ok := <-channel:
			if !ok {
				continue
			}

			var (
				filled bool
				price  int
			)

			if order.BuyOrSell {
				filled, price = matchOrderService.MatchBuyOrder(&order)
			} else {
				filled, price = matchOrderService.MatchSellOrder(&order)
			}

			if filled {
				fmt.Printf("Matched!! Price: %d\n", price)
			}
		default:
		}
	}
}

func sendOrderToChannel(buyOrSell bool, price int, count int) {
	for i := 0; i < count; i++ {
		order := domain.Order{
			BuyOrSell: buyOrSell,
			Quantity:  count,
			Price:     price, // = 0 aka marketing price
		}
		// put into channel, the channel is thread safe
		channel <- order
	}
}
