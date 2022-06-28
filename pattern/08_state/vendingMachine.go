package _8_state

import "fmt"

/*
Состояние — это поведенческий паттерн, позволяющий динамически изменять поведение объекта при смене его состояния.
Поведения, зависящие от состояния, переезжают в отдельные классы.
Первоначальный класс хранит ссылку на один из таких объектов-состояний и делегирует ему работу.
Концентрирует в одном месте код, связанный с определённым состоянием.


Минусы:
Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

// Общий интерфейс состояний
type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(money int) error
	dispenseItem() error
}

// Структура-контекст состояний
type vendingMachine struct {
	hasItem       state
	itemRequested state
	hasMoney      state
	noItem        state

	currentState state

	itemCount int
	itemPrice int
}

// Создание объекта vendingMachine с определенным состоянием (hasItemState) и параметрами
func newVendingMachine(itemCount, itemPrice int) *vendingMachine {
	v := &vendingMachine{
		itemCount: itemCount,
		itemPrice: itemPrice,
	}
	hasItemState := &hasItemState{
		vendingMachine: v,
	}
	itemRequestedState := &itemRequestedState{
		vendingMachine: v,
	}
	hasMoneyState := &hasMoneyState{
		vendingMachine: v,
	}
	noItemState := &noItemState{
		vendingMachine: v,
	}

	v.setState(hasItemState)
	v.hasItem = hasItemState
	v.itemRequested = itemRequestedState
	v.hasMoney = hasMoneyState
	v.noItem = noItemState
	return v
}

// Методы меняющие поведение в зависимости от состояния (currentState)

func (v *vendingMachine) requestItem() error {
	return v.currentState.requestItem()
}

func (v *vendingMachine) addItem(count int) error {
	return v.currentState.addItem(count)
}

func (v *vendingMachine) insertMoney(money int) error {
	return v.currentState.insertMoney(money)
}

func (v *vendingMachine) dispenseItem() error {
	return v.currentState.dispenseItem()
}

// Изменение состояния
func (v *vendingMachine) setState(s state) {
	v.currentState = s
}

// Увеличение товара в vendingMachine
func (v *vendingMachine) incrementItemCount(count int) {
	fmt.Printf("Adding %d items\n", count)
	v.itemCount = v.itemCount + count
}
