package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Состояние - поведенческий паттерн проектирования. Позволяет объектам менять свое поведение в зависимости
от состояния, извне кажется, что объект поменял свой тип.

Плюсы:
+ Избавления от множества больших условных операторов, машины состояния
+ Концентрирует в одном месте код связанный с определенным состояниям
+ Упрощает код контекста

Минусы:
- Если состояний мало и они редко меняются, неоправданно усложняет код
*/

//пример использование состояние
func exampleState() {
	water := newWater()  //инициализируем тип water
	_ = water.useWater() //попробуем выпить
	water.freezeWater()  //заморозим
	_ = water.useWater() //попробуем выпить еще раз
}

//тип water выступает в роли контекста, имеет поле для хранения объектов состояний
type water struct {
	//вода имеет два состояние - лед, жидкость
	ice    state
	liquid state

	currentState state
	//в зависимости от состояния воду можно пить или нет
	drink bool
}

//создаем тип water через конструктор
func newWater() (w *water) {
	w = &water{drink: true}

	liquidState := &liquidState{water: w}
	iceState := &iceState{water: w}

	w.liquid = liquidState
	w.ice = iceState

	w.SetState(liquidState)

	return w
}

//метод для изменения значения поля объектов состояний в контексте
func (w *water) SetState(st state) {
	w.currentState = st
}

//методы для типа water (использовать, заморозить, разморозить)
func (w *water) useWater() error {
	return w.currentState.useWater()
}
func (w *water) freezeWater() {
	w.currentState.freezeWater()
}
func (w *water) defrostWater() {
	w.currentState.defrostWater()
}

//общий интерфейс для состояний
type state interface {
	useWater() error
	freezeWater()
	defrostWater()
}

//типы конкретных состояний. имеет поле со ссылкой на конкретный тип
type iceState struct {
	water *water
}

func (i *iceState) useWater() error {
	return fmt.Errorf("Воду пить нельзя это кусок льда!\n Разморозьте ее!")
}
func (i *iceState) freezeWater() {
	fmt.Println("Вода уже заморожена!")
}

//состояние может менять контекст
func (i *iceState) defrostWater() {
	fmt.Println("Вода размораживается..")
	i.water.SetState(i.water.liquid)
	i.water.drink = true
}

//второй конкретный тип состояния
type liquidState struct {
	water *water
}

func (l *liquidState) useWater() error {
	fmt.Println("Пьем воду..")
	return nil
}
func (l *liquidState) freezeWater() {
	fmt.Println("Вода замораживается..")
	l.water.drink = false
	l.water.SetState(l.water.ice)

}
func (l *liquidState) defrostWater() {
	fmt.Println("Вода уже разморожена!")
}
