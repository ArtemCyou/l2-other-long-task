package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного
примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Строитель - порождающий паттерн программирования. Позволяет создавать сложные объекты пошагово.
Используя идин и тот же код строительства для получения разных продуктов.
*/

func builderPattern() {
	xiaomiBuilder := getBuilder("xiaomi") //создаем строителя который создаст продукт (телефон сяоми)

	//чтобы ни идти по шагам используем тип директор который знает как работать со строителем
	dir := newDirector(xiaomiBuilder)
	xiaomiPhone := dir.buildPhone() //получаем готовый продукт
	fmt.Println(xiaomiPhone)
}

//инициализирую наш будущий продукт, возможны разные его вариации
type phone struct {
	brand   string
	battery string
	price   float64
	//...
}

//код работает со строителями через общий интерфейс
type iBuilder interface {
	setBrand()
	setBattery()
	setPrice()
	getPhone() phone
}

//строитель который создает продукт продукт(телефон сяоми)
type xiaomiBuilder struct {
	brand   string
	battery string
	price   float64
}

//создаю строителя
func newXiaomiBuilder() *xiaomiBuilder {
	return &xiaomiBuilder{}
}

//пошагово создаю продукт с нужными для меня характеристиками
func (x *xiaomiBuilder) setBrand() {
	x.brand = "Xiaomi"
}
func (x *xiaomiBuilder) setBattery() {
	x.battery = "BIG"
}
func (x *xiaomiBuilder) setPrice() {
	x.price = 500
}

//возвращаю готовый продукт
func (x *xiaomiBuilder) getPhone() phone {
	return phone{
		brand:   x.brand,
		battery: x.battery,
		price:   x.price,
	}
}

//строитель который создает продукт продукт(телефон айфон)
type iphoneBuilder struct {
	brand   string
	battery string
	price   float64
}

//реализация данного строителя опущена...

//создадим тип директор, он знает как работать со строителем и автоматизирует нам работу с ним!
type director struct {
	builder iBuilder
}

func newDirector(b iBuilder) *director {
	return &director{builder: b}
}
//если нам потребуется поменять тип строителя используем сеттер
func (d *director) setBuilder(b iBuilder)  {
	d.builder = b
}

//инструкция для директора как работать со строителем и получить готовый продукт
func (d *director) buildPhone() phone {
	d.builder.setBattery()
	d.builder.getPhone()
	d.builder.setBrand()
	return d.builder.getPhone()
}

//функцию которая вернет нужный тип строителя
func getBuilder(builderType string) iBuilder {
	if builderType == "xiaomi" {
		return &xiaomiBuilder{}
	}

	if builderType == "iphone" {
		//return &iphoneBuilder{}
	}
	return nil
}
