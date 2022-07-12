package main

import (
	"bufio"
	"fmt"
	"io"
	"path/filepath"

	"log"
	"net/http"
	"net/url"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/
const (
	raw = iota
	good
	cancel
)

func main() {
	var f string                   //переменная для сохранения с stdin
	pages := make(map[url.URL]int) //мапка которая будет отслеживать обработана ссылка на сайт или нет

	//в бесконечном цикле получаем из stdin ссылки на сайт
	for {
		fmt.Print("Введите ссылку на сайт: ")
		scaner := bufio.NewScanner(os.Stdin)
		if scaner.Scan() {
			f = scaner.Text()
		}

		urlPars, err := url.Parse(f)

		if err != nil {
			log.Fatal(err)
		}

		pages[*urlPars] = raw

		for pageUrl, pageStatus := range pages {

			fmt.Println("Сайт:", pageUrl.String(), pageStatus, "\"0-обрабатывается, 1-сохранен\"")
			var site []byte

			if pageStatus == raw {

				site, pages[pageUrl] = checkUrlGetPage(pageUrl) //передаю юрл функции где проверим сайт и прочитаем, если он отвечает на запрос
				if pages[pageUrl] == good {
					err = saveSite(site, pageUrl.Hostname()) //сохраняем сайт
					if err != nil {
						log.Fatal(err) //завершаем программу и выводим ошибку, если не получилось сохранить сайт
					}
				}
			}
		}
	}
}

func saveSite(site []byte, url string) error {
	myDir, err := os.Getwd()                           //узнаем текущую директорию
	savePath := filepath.Join(myDir, "downloads", url) //создаем путь будущей директории

	err = os.MkdirAll(savePath, os.ModePerm) //создаем новую директорию для сохранения сайта
	if err != nil {
		fmt.Println("mkdir not found")
		return err
	}
	err = os.Chdir(savePath) //переходим в новую директорию
	if err != nil {
		fmt.Println("chmod not found")
		return err
	}

	//создаем файл и сохраняем в него сайт
	save, err := os.Create("site.html")
	defer save.Close()
	_, err = save.Write(site)
	if err != nil {
		return err
	}

	//возвращаемся обратно в изначальную директорию
	err = os.Chdir(myDir)
	if err != nil {
		return err
	}
	fmt.Println("Сайт сохранен")
	return nil
}

func checkUrlGetPage(pageUrl url.URL) ([]byte, int) {
	var site []byte
	response, err := http.Get(pageUrl.String())

	var errResp *http.Response = new(http.Response) //инициализирую указатель, чтобы не поймать панику в 105й строке
	if err != nil || response.StatusCode != 200 {
		fmt.Println("no get http", errResp)
		return nil, cancel
	}

	//response.Body.Read(site)
	site, err = io.ReadAll(response.Body)

	if err != nil {
		fmt.Println("no read body")
		return nil, cancel
	}
	defer response.Body.Close()

	return site, good
}
