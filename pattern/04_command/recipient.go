package _4_command

import "fmt"

type device interface {
	on()
	off()
}

// Получатель команд, реализует поведение device для связи с command
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}
