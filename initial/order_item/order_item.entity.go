package order_item

import (
	product "track-bootcamp/initial/product"
)

type OrderItem struct {
	Product product.Product
	Quantity int
}