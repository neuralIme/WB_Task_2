package _5_сhain_of_сommand

/*
Цепочка обязанностей — это поведенческий паттерн, позволяющий передавать запрос
по цепочке потенциальных обработчиков, пока один из них не обработает запрос

Избавляет от жёсткой привязки отправителя запроса к его получателю,
позволяя выстраивать цепь из различных обработчиков динамически.


Минусы:
Запрос может остаться никем не обработанным.
*/

// Интерфес department имеет два метода: один для связи между обработчиками, другой для обработки данных
type department interface {
	execute(*patient)
	setNext(department)
}

// Структура объекта
type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}
