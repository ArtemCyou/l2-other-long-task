package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Стратегия - поведенческий паттерн, позволяет определить схожие алгоритмы и каждый из них поместить в
собственный класс, после чего алгоритмы можно взаимозаменять прямо во время выполнения программы.

Плюсы:
+ Замена алгоритмов на лету
+ Изоляция кода алгоритмов от остальных типов
+ Реализует принцип открытости/закрытости

Минусы:
- Усложняет код за счет дополнительных типов
- Клиент должен знать в чем разница между стратегиями, чтобы выбрать нужный алгоритм
*/

//пример использования паттерна Стартегия
func exampleStrategy() {
	//инициализирую конкретную стратегию
	concreteStrategyAdd := &concreteStrategyAdd{}
	//создаю контекст и передаю стратегию
	context := newContext(concreteStrategyAdd, 2, 2)
	//запусскаю алгоритм
	context.execute()
	fmt.Println(context.number)

	//инициализирую новую стратегию и повторяю процедуру
	concreteStrategyMultiply := &concreteStrategyMultiply{}
	context.setStrategy(concreteStrategyMultiply)
	context.execute()
	fmt.Println(context.number)
}

	//контекст хранит ссылку на стратегию работает с ними через общий интерейс
	type sContext struct {
		strategy strategy
		a, b     int
		number   int
	}

	//конструктор для контекста
	func newContext(str	strategy, a, b int) *sContext{
		return &sContext{strategy: str,
		a:      a,
		b:      b,
		number: 0}
	}

	//передаем контексту определенную стратегию
	func(s *sContext) setStrategy(str strategy) {
		s.strategy = str
	}

	//вызываем алгоритм выполнения
	func(s *sContext) execute()	{
		s.strategy.execute(s)
	}

	//все стратегии имеют общий интерфейс
	type strategy interface {
		execute(sCon *sContext)
	}

	//конкретная стратегия (прибавление двух числе)
	type concreteStrategyAdd struct {
	}

	func(c *concreteStrategyAdd) execute(sCon * sContext){
		sCon.number = sCon.a + sCon.b
	}

	//конкретная стратегия (умножение двух числе)
	type concreteStrategyMultiply struct {
	}

	func(c *concreteStrategyMultiply) execute(sCon * sContext){
		sCon.number = sCon.a * sCon.b
	}
