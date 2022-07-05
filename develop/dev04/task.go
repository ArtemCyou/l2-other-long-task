package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
var mapAnagramm = make(map[string]*[]string)

func main() {
	arr := &[]string{"Пятак", "пятка", "Тяпка", "листок", "слиток", "столик", "отвар", "автор", "автор", "товар"}

	findAnagramm(arr)
}

func findAnagramm(arrG *[]string) *map[string]*[]string {
	arr := *arrG
	mapTemp := make(map[string][]string)
	//перебираем все слова в слайсе
	for _, val := range arr {
		//если слово короче 2х букв пропускаем его
		if len(val) <= 1 {
			continue
		}
		//приводим слово к нижнему регистру
		v := strings.ToLower(val)

		word := strings.Split(v, "")
		sort.Strings(word)

		key := strings.Join(word, "") //создаем ключ для анаграммы
		//проверяем есть ли такой ключ в мапе если нет, то создаем и добавляем слово в значение мапки
		if _, ok := mapTemp[key]; !ok {
			mapTemp[key] = make([]string,0)
			mapTemp[key] = append(mapTemp[key], v)
		} else {
			if searchVal(v, mapTemp) {
				mapTemp[key] = append(mapTemp[key], v)
			}
		}
	}

	//помещаем все значения в мапу со ссылкой на слайс, с ключом - первое слово из множества
	for _, val := range mapTemp{
		keyN := val[0]
		mapAnagramm[keyN] = new([]string)
		for _, v := range val {
			*mapAnagramm[keyN] = append(*mapAnagramm[keyN], v)
		}
	}

	//сортируем значения и выводим результат
	for i := range mapAnagramm {
		sort.Strings(*mapAnagramm[i])
		fmt.Println(*mapAnagramm[i])
	}

	return &mapAnagramm
}

//проверяем есть ли уже данное слово в значении мапке
func searchVal(value string, m map[string][]string) bool {
 for _, val := range m {
 	for _, v := range val{
 		if v == value {
			return false
		}
	}
 }
 return true
}
