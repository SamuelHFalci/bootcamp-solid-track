package order

import (
	"fmt"
	"log"
	paypal "v1/paypal"
)

type OrderService struct{
	
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (o *OrderService) Pay(order Order) {
	order.CalculateTotal()
	switch order.PaymentMethod {
	case "credit_card":
		fmt.Printf("Paying with credit card: %2.f", order.Total)
	case "paypal":
		paypalService := paypal.NewPaypalService()
		paypalService.Pay(order.Total)
	default:
		log.Fatal("Unknown payment method")
	}
}
