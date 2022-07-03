package main

import (

	"errors"
	"fmt"
	"strconv"

	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var errString = errors.New("некорректная строка")

const esc = '\\'

func main() {
	a := "qwe\\45" //"a4bc2d5e" => "aaaabccddddde" aaaabccddddde

	val, err := Unpack(a)
	fmt.Println(val, err)
	//- qwe\4\5 => qwe45 (*)
	//- qwe\45 => qwe44444 (*)
	//- qwe\\5 => qwe\\\\\ (*)
}

func Unpack(str string) (string, error) {
	var temp rune
	var res []rune
	var num int
	var escBool bool //флаг для escape последовательностей

	for i, val := range str {
		//проверяем флаг на наличие последовательностей
		if !escBool{
			//проверяем является ли текущий символ цифрой или escape последовательностью
			if unicode.IsDigit(val) || val == esc {
				//если цифра идет первой в очереди возвращаем ошибку
				if i == 0 {
					return "", errString
				}
				if val == esc {
					escBool = true
					continue
				}
				//преобразовываем руну в цифру
				num, _ = strconv.Atoi(string(val))
				//добавляем предыдущую руну в результирующий срез num-1 раз
				for i = 1; i < num; i++ {
					res = append(res, temp)
				}
				continue // переходим к следующей последовательности
			}
		}

		//если текущая руна не цифра и не escape последовательность добавляем в результирующий срез и
		//во временную переменную
		escBool = false
		res = append(res, val)
		temp = val
	}

	return string(res), nil
}
