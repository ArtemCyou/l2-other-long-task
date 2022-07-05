package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы []
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов []

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	//устанавливаем флаги запуска
	argK := flag.Int("k", 0, "указание колонки для сортировки")
	argSortNum := flag.Bool("n", false, "сортировать по числовому значению")
	argSortRev := flag.Bool("r", false, "сортировать в обратном порядке")
	argNoRept := flag.Bool("u", false, "не выводить повторяющиеся строки")
	argC := flag.Bool("c", false, "проверяет отсортированы ли данные")
	argM := flag.Bool("M", false, "сортировать по названию месяца")
	flag.Parse()


	//читаем файл по указанному пути, если путь не указали выводим ошибку
	if flag.Arg(0) == ""{
		fmt.Println("введите путь до файла")
		os.Exit(1)
	}
	//читаем файл
	fData := readF(flag.Arg(0))

	if *argK > 1 {
		*argK--
	}

	switch true {
	case *argSortNum: //n — сортировать по числовому значению
		sortInt(fData, *argK, *argSortRev)
	case *argNoRept: //u - не выводить повторяющиеся строки
		newFile := noRept(fData)
		fmt.Println(newFile)
	case *argC: //c — проверять отсортированы ли данные
		sortTrue(fData)
	case *argM: //M — сортировать по названию месяца
		sortMonth(fData, *argK, *argSortRev)

	default: //обычная сортировка, если есть флаг -r до будет обратная сортировка
		newFile := sortFun(fData, *argK, *argSortRev)
		fmt.Println(newFile)
	}

	fmt.Println(fData)
}

func readF(path string) [][]string {
	//откроем файл, через defer закроем его
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lineF [][]string //промежуточный срез в который сохраним данные файла

	//построчно прочитаем указанный файл и сохраним в срез
	scFile := bufio.NewScanner(file)
	for scFile.Scan() {
		line := strings.Split(scFile.Text(), " ")
		lineF = append(lineF, line)
	}

	return lineF
}

//сортировка по числовому значению
func sortInt(file [][]string, k int, revers bool) {

	sort.Slice(file, func(i, j int) bool {
		a, _ := strconv.ParseFloat(getFileElem(file, i, k), 64)
		b, _ := strconv.ParseFloat(getFileElem(file, j, k), 64)

		if revers {
			return a > b
		}
		return a < b
	})
}

//проверка отсортированы ли данные
func sortTrue(file [][]string) {
	res := sort.SliceIsSorted(file, func(i, j int) bool {
		return file[i][0] < file[j][0]
	})

	if res {
		fmt.Println("Данные отсортированы")
		return
	}
	fmt.Println("Данные не отсортированы")
}

//сортировка значений так-же и в обратном порядке
func sortFun(file [][]string, k int, revers bool) [][]string {
	//sort.Sort(sort.Reverse(sort.StringSlice(file))) не работает с двумерным массивом
	if revers {
		sort.Slice(file, func(i, j int) bool {
			return getFileElem(file, i, k) > getFileElem(file, j, k)
		})
	} else {
		sort.Slice(file, func(i, j int) bool {
			return getFileElem(file, i, k) < getFileElem(file, j, k)
		})
	}
	return file
}

func sortMonth(file [][]string, k int, revers bool) {
	if revers {
		sort.Slice(file, func(i, j int) bool {
			return getFileMonth(getFileElem(file, j, k)).Before(getFileMonth(getFileElem(file, i, k)))
		})
	} else {
		sort.Slice(file, func(i, j int) bool {
			return getFileMonth(getFileElem(file, i, k)).Before(getFileMonth(getFileElem(file, j, k)))
		})
	}
}

//возвращает элемент из слайса под индексом k
func getFileElem(file [][]string, i, k int) string {
	if k < len(file[i]) {
		return file[i][k]
	}
	return ""
}

func getFileMonth(month string) time.Time {
	if t, err := time.Parse("март", month); err == nil {
		return t
	}
	if t, err := time.Parse("февраль", month); err == nil {
		return t
	}
	if t, err := time.Parse("январь", month); err == nil {
		return t
	}
	if t, err := time.Parse("1", month); err == nil {
		return t
	}
	if t, err := time.Parse("01", month); err == nil {
		return t
	}
	return time.Time{}
}

//убираем повторяющиеся строки
func noRept(file [][]string) [][]string {
	mapBool := make(map[string]bool)
	res := make([][]string, len(file))

	for i, vSl := range file {
		for _, v := range vSl {
			if _, ok := mapBool[v]; !ok {
				mapBool[v] = true
				res[i] = append(res[i], v)
			}
		}
	}
	inCmd(res)
	return res
}

//функция вывода результата в командной строке
func inCmd(file [][]string) {
	for _, sl := range file {
		fmt.Println(strings.Join(sl, " "))
	}
}
