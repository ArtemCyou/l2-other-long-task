package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	//устанавливаю флаги для работы утилиты
	argF := flag.String("f", "", "f - \"fields\" - выбрать поля (колонки)")
	argD := flag.String("d", " ", "d - \"delimiter\" - использовать другой разделитель")
	argS := flag.Bool("s", false, "s - \"separated\" - только строки с разделителем")
	flag.Parse()
	var words []string

	//принимаем строки в stdin в бесконечном цикле
	for {
		scanner := bufio.NewScanner(os.Stdin)
		var txt string
		if scanner.Scan() {
			txt = scanner.Text()
		}
		fmt.Printf("%#v\n", txt) //goto не забыть удалить

		if !*argS {
			//учитываем флаг d "delimiter" - использовать другой разделитель
			words = strings.Split(txt, *argD) //полученную строку из stdin разбиваем и сохраняем в слайс
			}

		switch true {
		case *argF != "": //f - "fields" - выбрать поля (колонки)
			//преобразуем строку с номерами колонок в числовой слайс
			columns := splitColums(*argF)
			//перебираем полученную строку и выводим только нужные колонки
			for i, val := range words {
				for _, v := range columns {
					if i == v-1 {
						fmt.Println(val)
					}
				}
			}

		case *argS:
			wordsSep := onlySep(*argD, txt)
			fmt.Println(wordsSep)

		default: //если ключи f не активирован то разбиваем по разделителю (TAB) на колонки и выводим все
			for _, val := range words {
				fmt.Println(val)
			}
		}
	}
}

//строку с номерами колонок превращает в числовой слайс
func splitColums(colString string) (columns []int) {
	for _, val := range colString {
		num, err := strconv.Atoi(string(val))
		if err != nil {
			continue
		}
		columns = append(columns, num)
	}
	return columns
}

//возвращаем левую и правую колонку от разделителя
func onlySep(sep string, txt string) []string {
	if sep == " " {
	 	words := strings.Split(txt, " ")
	 	return words
	}
	onlyS := sep
	var onlyWordsSep []string

	words := strings.Split(txt, sep)
	for i, val := range words{
		if val == onlyS{
			onlyWordsSep = append(onlyWordsSep, words[i-1], words[i], words[i+1])
		}
	}

	return words
}
