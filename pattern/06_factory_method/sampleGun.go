package _6_factory_method

/*
Фабричный метод — это порождающий паттерн проектирования,
который решает проблему создания различных продуктов,
без указания конкретных классов продуктов.

Минусы:
Может привести к созданию больших параллельных иерархий классов,
так как для каждого класса продукта надо создать свой подкласс создателя
*/

// Общий интерфейс, который реализуют все создаваемые объекты
type iGun interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Шаблон создаваемых объектов, реализует интерфейс iGun
type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}
