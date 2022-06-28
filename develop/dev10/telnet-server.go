package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// serverTCP запускает сервер на порту 8081
// полученные строки от клиента возвращает обратно, увеличивая регистр
func serverTCP() {

	fmt.Println("Запуск сервера...")

	ln, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}

	for {

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print("Сообщение доставлено:", message)

		newMessage := strings.ToUpper(message)

		_, err = conn.Write([]byte(newMessage + "\n"))
		if err != nil {
			log.Println("отправка не удалась: клиент не доступен")
		}
	}
}
