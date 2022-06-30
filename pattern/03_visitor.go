package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Посетитель - поведенческий паттерн, позволяет добавить новый метод для целой иерархии типов не изменяя
код этих типов.

Плюсы:
+ Проще добавлять операции для разных типов в проекте
+ Объединяет родственные операции в одном типе

Минусы:
- Может привести к нарушению инкапсуляции элементов
- Если в проекте часто добавляются новые типы применение паттерна не оправдано
*/

//пример использования паттерна
func visitorForFood()  {
	//инициализируем разный тип еды
	iceCream := &iceCream{}
	coffee := &coffee{}

	//инициализируем посетителя
	tasteVisitor := tasteVisitor{}

	//вызываем метод который попробует еду на вкус не зависимо от типа еды
	iceCream.accept(tasteVisitor)
	coffee.accept(tasteVisitor)
}

//общий интерфейс для еды
type food interface {
	//... other method
	accept(v visitor)
}

//каждый новый тип еды реализует общий интерфейс для еды
type iceCream struct {}

func (i *iceCream) accept(v visitor) {
	v.visitorForIceCream(i)
}

type coffee struct {}

func (c *coffee) accept(v visitor) {
	v.visitorForCoffee(c)
}

type borscht struct {}

func (b *borscht) accept(v visitor) {
	v.visitorForBorscht(b)
}

//интерфейс посетителя, он знает о каждом типе еды
type visitor interface {
	visitorForIceCream(i *iceCream)
	visitorForCoffee(c *coffee)
	visitorForBorscht(b *borscht)
}

//наш конкретный посетитель который пробует еду на вкус
type tasteVisitor struct {
}

//реализация методов для каждого конкретного типа еды
func (t tasteVisitor) visitorForIceCream(i *iceCream) {
	//.. other code
	fmt.Println("Вкусное мороженое!")
}
func (t tasteVisitor) visitorForCoffee(c *coffee) {
	fmt.Println("Вкусный кофий!")
}
func (t tasteVisitor) visitorForBorscht(b *borscht) {
	fmt.Println("Вкусный борщстч!")
}
