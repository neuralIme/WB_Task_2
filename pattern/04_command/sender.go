package _4_command

/*
 button - структура кнопки-отправителя сигнала для включения тв, наследник команды
*/

type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}
