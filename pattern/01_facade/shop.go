package _1_facade

import "time"

// Магазин
type Shop struct {
	products []Product
}

// Товар магазина
type Product struct {
	name  string
	price float64
}

// Метод CheckProduct проверяет наличие товара в магазине
func (shop *Shop) CheckProduct(product string) (float64, bool) {
	time.Sleep(300 * time.Millisecond)
	for _, prod := range shop.products {
		if prod.name == product {
			return prod.price, true
		}
	}
	return 0, false
}
