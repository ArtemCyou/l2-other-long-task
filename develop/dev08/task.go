package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	var command string
	//в цикле ждем ввод команды и отправляем их exec
	for {
		fmt.Print("Введите команду: ")
		scaner := bufio.NewScanner(os.Stdin)

		if scaner.Scan() {
			command = scaner.Text()
		}
		if command == "/quit" { //проверяем команду выхода
			os.Exit(0)
		}
		out, err := exeC(command)
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Printf("Результат: %s", out)
	}

}

func exeC(command string) (out []byte, err error) {

	cmd := exec.Command("bash") //запускаем баш

	//получаем канал который подключен к стандартному вводу команды
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	//создаем входной канал
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, command)
	}()

	out, err = cmd.CombinedOutput() //получаем результат выполненной команды
	if err != nil {
		log.Fatal(err)
	}
	return out, err
}
