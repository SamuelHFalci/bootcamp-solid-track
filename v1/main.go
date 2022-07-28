package main

import (
	order "v1/order"
	orderItem "v1/order_item"
	product "v1/product"
)

func main() {
	o := order.Order{
		Id: 1,
		CustomerId: 1,
		OrderItems: []orderItem.OrderItem{
			{
				Product: product.Product{
					Name: "Product 1",
					Price: 10,
				},
				Quantity: 1,
			},
			{
				Product: product.Product{
					Name: "Product 2",
					Price: 20,
				},
				Quantity: 2,
			},
		},
		Total: 0,
		PaymentMethod: "paypal",
	}

	orderService := order.NewOrderService()
	orderService.Pay(o)
}