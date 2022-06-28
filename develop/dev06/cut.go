package main

import (
	"bufio"
	"os"
	"strings"
)

type iCut interface {
	redact(cmd []string)
}

// Cut структура объекта, управляющего методами
type Cut struct {
	err    string
	method iCut
}

// setMethod устанавливает метод для вывода данных
func (c *Cut) setMethod(cmd []string) {
	c.err = "некорректный ввод"
	switch cmd[0] {
	case "-f":
		c.method = f
		f.redact(cmd[1:])
	case "-d":
		c.method = d
		d.redact(cmd[1:])
	case "-s":
		c.method = s
		s.redact(cmd[2:])
	default:
		println(c.err)
	}
}

var (
	f = &fields{}
	d = &delimiter{}
	s = &separated{}
)

var data [][]string

func main() {

	var (
		cut     = Cut{}
		scanner = bufio.NewScanner(os.Stdin)
	)

	for scanner.Scan() {
		str := strings.Split(scanner.Text(), " ")
		switch str[0] {
		case "cut":
			if len(str) > 1 {
				cut.setMethod(str[1:])
			} else {
				continue
			}
		case " ", "\n", "", "\t":
			continue
		default:
			data = append(data, str)
		}
	}
}
