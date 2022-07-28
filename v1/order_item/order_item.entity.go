package order_item

import (
	product "v1/product"
)

type OrderItem struct {
	Product product.Product
	Quantity int
}