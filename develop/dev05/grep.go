package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// iGrep общий интерфейс объектов
type iGrep interface {
	printResult(args []string)
}

// Grep структура для связи методов
type Grep struct {
	err    string
	method iGrep
}

// readFile читает данные из файла
func (g *Grep) readFile(fileName string) {
	g.err = "[readFile] не удалось получить данные из файла"
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(g.err, err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data = append(data, strings.TrimSuffix(scanner.Text(), "\n"))
	}
}

// distributor распределяет методы в зависимости от команды пользователя
func (g *Grep) distributor(args []string) {
	switch args[0] {
	case "-A":
		fmt.Println(args)
		g.method = after
		after.printResult(args[1:])
	case "-B":
		g.method = before
		before.printResult(args[1:])
	case "-C":
		g.method = context
		context.printResult(args[1:])
	case "-c":
		g.method = count
		count.printResult(args[1:])
	case "-i":
		g.method = ignoreCase
		ignoreCase.printResult(args[1:])
	case "-v":
		g.method = invert
		invert.printResult(args[1:])
	case "-F":
		g.method = fixed
		fixed.printResult(args[1:])
	case "-n":
		g.method = lineNum
		lineNum.printResult(args[1:])
	default:
		g.method = search
		search.printResult(args)
	}

}

var (
	grep       = &Grep{}
	search     = &Search{}
	after      = &After{}
	before     = &Before{}
	context    = &Context{}
	count      = &Count{}
	ignoreCase = &IgnoreCase{}
	invert     = &Invert{}
	fixed      = &Fixed{}
	lineNum    = &LineNum{}
	data       []string
)

func main() {

	args := os.Args[1:]

	grep.readFile(args[len(args)-1])

	switch args[0] {
	case "-A", "-B", "-C", "-c", "-i", "-v", "-F", "-n":
		grep.distributor(args[:len(args)-1])
	default:
		if len(args) != 2 {
			log.Fatal(fmt.Sprintf("[main] некорректный ввод: %s", args))
		}
		grep.distributor(args[:len(args)-1])
	}
}
