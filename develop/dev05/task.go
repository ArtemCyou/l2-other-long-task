package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	//установка определенных настроек для запуска утилиты
	argA := flag.Int("A", 0, "\"after\" печатать +N строк после совпадения")
	argB := flag.Int("B", 0, "\"before\" печатать +N строк до совпадения")
	argC := flag.Int("C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	argCsmall := flag.Bool("c", false, "\"count\" (количество строк)")
	argIsmall := flag.Bool("i", false, "\"ignore-case\" (игнорировать регистр)")
	argVsmall := flag.Bool("v", false, "\"invert\" (вместо совпадения, исключать)")
	argF := flag.Bool("F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	argNsmall := flag.Bool("n", false, "\"line num\", печатать номер строки")

	flag.Parse()

	//задаем что ищем и где ищем
	strSearch, pathFile := flag.Arg(0), flag.Arg(1)
	log.Printf("str= %s, path=%s", strSearch, pathFile)
	dataF := readF(pathFile, *argIsmall) //предусмотрен ключ i

	if pathFile == "" || strSearch == "" {
		log.Println("Введите команду типа:\n \"искомая строка\" \"путь до файла для чтения\"")
		os.Exit(1)
	}
	//если активирован флаг i "ignore-case" приводим все строки к одному регистру
	if *argIsmall {
		strSearch = strings.ToLower(strSearch)
	}

	//ищем строку в слайсе и сохраняем ее индекс если есть совпадение
	//предусмотрен ключ v - "invert"
	var iSrchIndx []int
	for i, val := range dataF {
		if strings.Contains(val, strSearch) && !*argVsmall {
			iSrchIndx = append(iSrchIndx, i)
		} else if !strings.Contains(val, strSearch) && *argVsmall {
			iSrchIndx = append(iSrchIndx, i)
		}
	}

	switch true {
	case *argA > 0: //A - выводим в консоль N+ строк после совпадения
		for _, ind := range iSrchIndx {
			for i := 0; i < *argA; i++ {
				if ind+i > len(dataF) {
					fmt.Println(dataF[ind+i])
				}
			}
		}

	case *argB > 0: //B -  выводим в консоль N+ строк до  совпадения
		for _, ind := range iSrchIndx {
			for i := 0; i < *argB; i++ {
				if ind-*argB+i <= 0 {
					idx := *argB + i
					fmt.Println(dataF[idx])
				}
			}
		}

	case *argC > 0: //C -  выводим в консоль N+ строк до и после совпадения
		for _, ind := range iSrchIndx {
			for i := 0; i < *argC; i++ {
				if ind+i > len(dataF) {
					fmt.Println(dataF[ind+i])
				}
				if ind-*argC+i <= 0 {
					idx := *argC + i
					fmt.Println(dataF[idx])
				}
			}
		}

	case *argCsmall: //с - выводим "count" (количество строк)
		fmt.Println(len(iSrchIndx))

	case *argNsmall: //n - "line num", печатать номер строки
		for _, lNum := range iSrchIndx { // lnum+1 так как строка в слайсе имеет счет с 0
			fmt.Println(lNum + 1)
		}

	case *argF: //F - "fixed", точное совпадение со строкой, не паттерн
		newIsearchInd := hardCont(dataF, strSearch)
		for _, ind := range newIsearchInd {
			fmt.Println(dataF[ind])
		}

	default: //выводим результат по умолчанию, все строки где есть совпадение
		for _, ind := range iSrchIndx {
			fmt.Println(dataF[ind])
		}
	}
}

//для работы флага F - "fixed", точное совпадение со строкой, не паттерн
func hardCont(dataF []string, strSearch string) (newIsearchInd []int) {
	for i, val := range dataF {
		if val == strSearch {
			newIsearchInd = append(newIsearchInd, i)
		}
	}
	return newIsearchInd
}

//читаем файл возвращаем слайс
func readF(path string, i bool) (file []string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	nSc := bufio.NewScanner(f)
	for nSc.Scan() {
		if i { //если активирован флаг i "ignore-case" приводим все строки к одному регистру
			file = append(file, strings.ToLower(nSc.Text()))
		} else {
			file = append(file, nSc.Text())
		}
	}

	return file
}
