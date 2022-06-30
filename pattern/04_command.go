package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Команда - поведенческий паттерн, позволяет упаковать запрос в отдельный объект

Плюсы:
+ Можно реализовать отложенный запуск операций
+ Можно реализовать отмену и повтор операции
+ Убирает зависимость между объектом который вызывает операцию и объектом который ее выполняет

Минусы:
- Усложняется код программы из-за появления большого количества дополнительных типов
*/

//пример использования паттерна команда
func onAndOffCommand()  {
	lp:= &lamp{} //инициализируем девайс лампа

	onCom := &onCommand{device: lp} //инициализируем команду которая включит девайс лампу
	offCom := &offCommand{device: lp}

	onBut := &button{command: onCom} //инициализируем кнопку и передаем ей команду
	onBut.press() //вызываем метод, который активирует команду

	offBut := &button{command: offCom}
	offBut.press()
}

//отправитель команды
type button struct {
	command command
}

//метод отправителя вызывающий команду
func (b *button) press() {
	b.command.execute()
}

//общий интерфейс команды
type command interface {
	execute()
}

//конкретная команда включающая лампу. реализует общий интерфейс команды
type onCommand struct {
	//имеет поле с получателем которым команда будет перенаправлять работу
	device device
}

//метод выключающий лампу
func (o *onCommand) execute() {
	o.device.off()
}

//конкретная команда выключающая лампу
type offCommand struct {
	device device
}

func (o *offCommand) execute() {
	o.device.on()
}

//интерфейс получателя
type device interface {
	on()
	off()
}

//конкретный получатель команды
type lamp struct {
	itWorked bool
}

//методы получателя команды
func (l *lamp) on() {
	l.itWorked = true
	fmt.Println("Лампа включена!")
}
func (l *lamp) off() {
	l.itWorked = false
	fmt.Println("Лампа выключена!")
}
