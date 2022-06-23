package pattern

import "fmt"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера
на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Применяется когда нужно обращаться со сложной подсистемой, реализует простой интерфейс для доступа к
определенной функциональности подсистемы.

Плюсы:
-Изолирует клиента от сложной подсистемы
-Снижает сложность программы
-Можно вывести код зависимый от внешней системы в единое место

Минусы:
-Может стать "Божественным объектом" привязанными ко всем классам программы
*/

//реализуем паттерн фасад
type searchFacade struct {
	query *searchInChrome
	clean *cleanHistoryChrome
}

//создаем searchFacade
func newSearchFacade() *searchFacade {
	searchFaccade := &searchFacade{
		query: nil,
		clean: nil,
	}
	return searchFaccade
}

//реализуем простой интерфейс для взаимодействия с подсистемами
func (s *searchFacade) searchAndClean(query string) (result string) {
res:= s.query.search(query)
s.clean.clean()
return res
}

//searchInChrome реализует подсистему №1
type searchInChrome struct {

}
func (s *searchInChrome)search(query string) string {
 //обрабатывает query возвращает результат
	return "результат по запросу: " + query
}

//cleanHistoryChrome реализует подсистему №2
type cleanHistoryChrome struct {

}
func (c *cleanHistoryChrome) clean()   { //удаляем историю поиска
	fmt.Println("история успешно очищена")
}