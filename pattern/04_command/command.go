package _4_command

/*
Команда — это поведенческий паттерн, позволяющий заворачивать запросы или простые операции в отдельные объекты.
Это позволяет откладывать выполнение команд, выстраивать их в очереди, а также хранить историю и делать отмену.


Минусы:
Усложняет код программы из-за введения множества дополнительных классов.
*/

type command interface {
	execute()
}

// Команда для включения тв с соответствующим методом, наследник получателя
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// Команда для выключения тв с соответствующим методом, наследник получателя
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}
