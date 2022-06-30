package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/
/*
Цепочка вызовов - поведенческий паттерн, позволяет передавать запросы последовательно по цепочке обработчиков.

Плюсы:
+ Реализует принцип единственной обязанности
+ Реализует принцип открытости/закрытости
+ Уменьшает зависимость между клиентом и обработчиками

Минусы:
- Запрос может остаться не обработанным
*/

//пример использования паттерна Цепочка вызовов
func chainExample() {
	//инициализирую тип который будут обрабатывать обработчики
	car := &car{model: "nissan"}

	//инициализирую обработчики
	engine := &engine{}

	diagnostics := &diagnostics{}
	//после обработки обработчик передаст объект следующему обработчику по цепи
	diagnostics.setNext(engine)
	//первый обработчик в цеи
	diagnostics.execute(car)
}

//интерфейс обработчика имеет метод обработки запроса, и метод вызов следующего обработчика
type sto interface {
	execute(*car)
	setNext(sto) //сеттер чтобы задать обработчик
}

//конкретный обработчик имеет поле для следующего обработчика
type diagnostics struct {
	next sto
}

func (d *diagnostics) execute(c *car) {
	if c.diagnosticsOk {
		fmt.Println("Машина уже успешно прошла диагностику")
		d.next.execute(c)
		return
	}
	fmt.Println("Машина проходит диагностику")
	//other code
	c.diagnosticsOk = true
	d.next.execute(c)
}

func (d *diagnostics) setNext(next sto) {
	d.next = next
}

//еще один конкретный разработчик
type engine struct {
	next sto
}

func (e *engine) execute(c *car) {
	if c.engineOil {
		fmt.Println("В моторе свежее масло")
		return
	}
	fmt.Println("Меняем масло в моторе..")
	c.engineOil = true

}
func (e *engine) setNext(next sto) {
	e.next = next
}

//конкретный тип который обрабатывают обработчики в цепи
type car struct {
	model         string
	diagnosticsOk bool
	engineOil     bool
}
