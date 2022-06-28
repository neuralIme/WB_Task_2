package _1_facade

/*
Фасад – это структурный паттерн проектирования, который предоставляет простой интерфейс
к сложной системе классов, библиотеке или фреймворку. Изолирует клиентов от компонентов сложной подсистемы.


Однако данный паттерн рискует стать "божественным объектом", привязанным ко всем классам программы.
*/

import (
	"fmt"
)

// Клиент имеет карту банка
type Client struct {
	id    int
	money Card
}

// Метод Cell является Фасадом в данном примере. Сначала он проверяет наличие товара в магазине и его цену,
// используя метод магазина, после проверяет баланс карты банка, используя метод банка, после чего совершает продажу
func (shop *Shop) Cell(client Client, product string) {
	price, ok := shop.CheckProduct(product)
	if !ok {
		println("товар не найден")
		return
	}
	if price > client.money.CheckBalance() {
		println("недостаточно средств")
		return
	}
	client.money.balance -= price
	println(fmt.Sprintf("покупка совершена успешно!\nсписано -%f", price))
}
