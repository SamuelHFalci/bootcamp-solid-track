package order

import (
	orderItem "v1/order_item"
)

type Order struct {
	Id int64
	CustomerId int64
	OrderItems []orderItem.OrderItem
	Total float64
	PaymentMethod string
}

func (o *Order) CalculateTotal() {
	for _, item := range o.OrderItems {
		o.Total += item.Product.Price * float64(item.Quantity)
	}
}