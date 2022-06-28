package _7_strategy

import "fmt"

/*
Стратегия — это поведенческий паттерн, выносит набор алгоритмов в собственные классы и делает их взаимозаменимыми.
Другие объекты содержат ссылку на объект-стратегию и делегируют ей работу.
Программа может подменить этот объект другим, если требуется иной способ решения задачи.
Изолирует код и данные алгоритмов от остальных классов.

Минусы:
Усложняет программу за счёт дополнительных классов/структур.
Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/

// Объекты evictionAlgo являются стратегиями очищения кэша в данном примере
type evictionAlgo interface {
	evict(c *cache)
}

type fifo struct {
}

func (l *fifo) evict(c *cache) {
	fmt.Println("убирает запись, которая была создана раньше остальных")
}

type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("убирает запись, которая использовалась наиболее давно")
}

type lfu struct {
}

func (l *lfu) evict(c *cache) {
	fmt.Println("убирает запись, которая использовалась наименее часто")
}