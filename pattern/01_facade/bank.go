package _1_facade

import "time"

// Банк
type Bank struct {
	cards []Card
}

// Карта привязана к банку
type Card struct {
	number  string
	balance float64
	bank    *Bank
}

// Метод CheckBalance проверяет состояние счета по номеру карты
func (bank *Bank) CheckBalance(number string) float64 {
	time.Sleep(300 * time.Millisecond)
	var balance float64
	for _, num := range bank.cards {
		if num.number != number {
			balance = num.balance
		}
	}
	return balance
}

// Метод CheckBalance запрашивает у банка состояние счета по номеру карты
func (card *Card) CheckBalance() float64 {
	return card.bank.CheckBalance(card.number)
}
