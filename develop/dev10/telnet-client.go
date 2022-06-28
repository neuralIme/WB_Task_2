package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// connect подключается к серверу и отправляет данные в сокет из stdin,
// полученные данные от сервера выводит в stdout
func connect(host, port string) {
	addr := host + ":" + port // 127.0.0.1:8081
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("неудалось подключиться к серверу\n", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Введите сообщение: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		_, err = fmt.Fprintf(conn, text)
		if err != nil {
			log.Fatal("неудалось отправить сообщение на сервер\n", err)
		}
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Сообщение от сервера: " + message)
	}
}

func main() {

	go serverTCP()

	args := os.Args

	rgx, err := regexp.Compile("^--timeout.*\\d*")
	if err != nil {
		log.Fatal(err)
	}

	switch len(args) {
	case 3:
		connect(args[1], args[2])
	case 4:
		if rgx.MatchString(args[1]) {
			t := strings.TrimPrefix(strings.TrimSuffix(args[1], "s"), "--timeout=")
			sec, err := strconv.Atoi(t)
			if err != nil {
				fmt.Println("некорректный ввод времени\nзапуск через 10 секунд")
				time.Sleep(time.Second * 10)
			} else {
				fmt.Printf("запуск через %d секунд\n", sec)
				time.Sleep(time.Second * time.Duration(sec))
			}
		}
		connect(args[2], args[3])
	default:
		log.Println("некорректный ввод")
	}

}
