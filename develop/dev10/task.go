package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	timeout := flag.Duration("timeout", 1*time.Second, "таймаут на подключение к серверу")
	flag.Parse()

	//flag.Arg принимает аргументы п
	host := os.Args[2]
	port := os.Args[3]
	conAdress := fmt.Sprint(host, ":", port)

	//инициализирую канал для работы shut down
	killChan := make(chan os.Signal, 1)
	signal.Notify(killChan, os.Interrupt)

	//передаем время для таймаута, адрес подключения и канал
	clientTCP(*timeout, conAdress, killChan)

}

func clientTCP(timeout time.Duration, conAdress string, killChan <-chan os.Signal) {
	//инициализирую нужные переменные
	var flags = true
	var err error
	var conn net.Conn
	//инициализирую таймер таймаута
	timer := time.After(timeout)
	tm := time.Now()

	for flags {
		select {
		case <-timer:
			fmt.Println("Время ожидания подключения вышло!")
			fmt.Println("Ожидали:", time.Since(tm), conAdress)
			flags = false
			return
		default:
			conn, err = net.Dial("tcp", conAdress)
			if err != nil {
				time.Sleep(100 * time.Millisecond)
				continue
			}
			flags = false
			break
		}
	}

	//в отдельной горутине следим за сигналом системы для завершения программы
	go shutDown(killChan, conn)

	//в бесконечном цикле пересылаем данные
	for {
		//читаем из stdin
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Input: ")
		input , _ := reader.ReadString('\n')

		//отправляем в stdout
		fmt.Fprintf(conn, input +"\n")
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("#%s\n", err)
		}
		fmt.Print("Ответ: ", message)
		//conn.Close()
	}
}

func shutDown(killChan <-chan os.Signal, conn net.Conn) {
	select {
	case <-killChan:
		conn.Close()
	}
}
