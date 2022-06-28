package main

import (
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
)

// distributor распределяет команды по функциям
func distributor(cmd []string) error {
	switch cmd[0] {
	case "cd":
		return cd(cmd[1:])
	case "pwd":
		return pwd()
	case "echo":
		return echo(cmd[1:])
	case "kill":
		return kill(cmd[1:])
	case "ls":
		return ls()
	case "ps":
		return ps()
	case "exec":
		return exe(cmd[1:])
	case "fork":
		fork()
		return nil
	default:
		return errors.New("команда не найдена")
	}
}

func main() {

	scannner := bufio.NewScanner(os.Stdin)

	for scannner.Scan() {

		input := strings.Split(scannner.Text(), " ")

		if input[0] == "\\quit" {
			return
		}

		if err := distributor(input); err != nil {
			log.Println(err)
		}
	}
}
