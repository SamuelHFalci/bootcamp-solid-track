package main

import (
	order "track-bootcamp/initial/order"
	orderItem "track-bootcamp/initial/order_item"
	product "track-bootcamp/initial/product"
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
	
	// TODO: implementar o módulo de pagamento
	orderService := order.NewOrderService()
	orderService.Pay(o)
}