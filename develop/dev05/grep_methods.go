package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

// Search реализует интерфейс iGrep
type Search struct {
	err string
}

// printResult печатает совпадения
func (s *Search) printResult(args []string) {
	s.err = fmt.Sprintf("[Search] некорректный ввод: %s", args)
	reg, err := regexp.Compile(args[0])
	if err != nil {
		log.Fatal(s.err, "\n", err)
	}
	for _, str := range data {
		if reg.MatchString(str) {
			fmt.Println(str)
		}
	}
}

// After реализует интерфейс iGrep
type After struct {
	err string
}

// printResult печатает +N строк после совпадения
func (a *After) printResult(args []string) {
	a.err = fmt.Sprintf("[After] некорректный ввод: %s", args)
	if len(args) != 2 {
		log.Fatal(a.err)
	}
	reg, err := regexp.Compile(args[1])
	if err != nil {
		log.Fatal(a.err, "\n", err)
	}
	count, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(a.err, "\n", err)
	}
	for i := range data {
		if reg.MatchString(data[i]) {
			for j := i; j < i+count; j++ {
				if j < len(data) {
					fmt.Println(data[j])
				}
			}
		}
	}
}

// Before реализует интерфейс iGrep
type Before struct {
	err string
}

// printResult печатает +N строк до совпадения
func (b *Before) printResult(args []string) {
	b.err = fmt.Sprintf("[Before] некорректный ввод: %s", args)
	if len(args) != 2 {
		log.Fatal(b.err)
	}
	reg, err := regexp.Compile(args[1])
	if err != nil {
		log.Fatal(b.err, "\n", err)
	}
	count, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(b.err, "\n", err)
	}
	for i := range data {
		if reg.MatchString(data[i]) {
			for j := i - count; j < i; j++ {
				if j >= 0 {
					fmt.Println(data[j])
				}
			}
			fmt.Println(data[i])
		}
	}
}

// Context реализует интерфейс iGrep
type Context struct {
	err string
}

// printResult печатает +N строк вокруг совпадения
func (c *Context) printResult(args []string) {
	c.err = fmt.Sprintf("[Context] некорректный ввод: %s", args)
	if len(args) != 2 {
		log.Fatal(c.err)
	}
	reg, err := regexp.Compile(args[1])
	if err != nil {
		log.Fatal(c.err, "\n", err)
	}
	count, err := strconv.Atoi(args[0])
	if err != nil {
		log.Fatal(c.err, "\n", err)
	}
	for i := range data {
		if reg.MatchString(data[i]) {
			for j := i - count; j < i+count; j++ {
				if j >= 0 && j < len(data) {
					fmt.Println(data[j])
				}
			}
		}
	}
}

// Count реализует интерфейс iGrep
type Count struct {
	count int
	err   string
}

// printResult печатает количество совпадающих строк
func (c *Count) printResult(args []string) {
	c.err = fmt.Sprintf("[Count] некорректный ввод: %s", args)
	if len(args) != 1 {
		log.Fatal(c.err)
	}
	reg, err := regexp.Compile(args[0])
	if err != nil {
		log.Fatal(c.err, "\n", err)
	}
	for i := range data {
		if reg.MatchString(data[i]) {
			c.count++
		}
	}
	fmt.Println(c.count)
}

// IgnoreCase реализует интерфейс iGrep
type IgnoreCase struct {
	err string
}

// printResult игнорирует регистр
func (i *IgnoreCase) printResult(args []string) {
	args[len(args)-1] = fmt.Sprintf("(?is)%s", args[len(args)-1])
	grep := &Grep{}
	grep.distributor(args)
}

// Invert реализует интерфейс iGrep
type Invert struct {
	err string
}

// printResult печатает строки исключая совпадения
func (i *Invert) printResult(args []string) {
	i.err = fmt.Sprintf("[Invert] некорректный ввод: %s", args)
	reg, err := regexp.Compile(args[0])
	if err != nil {
		log.Fatal(i.err, "\n", err)
	}
	for _, str := range data {
		if reg.MatchString(str) {
			continue
		}
		fmt.Println(str)
	}
}

// Fixed реализует интерфейс iGrep
type Fixed struct {
	err string
}

// printResult печатает точное совпадение со строкой, не паттерн
func (f *Fixed) printResult(args []string) {
	args[len(args)-1] = fmt.Sprintf("^%s$", args[len(args)-1])
	grep := &Grep{}
	grep.distributor(args)
}

// LineNum реализует интерфейс iGrep
type LineNum struct {
	err string
}

// printResult печатает номер строки
func (l *LineNum) printResult(args []string) {
	l.err = fmt.Sprintf("[LineNum] некорректный ввод: %s", args)
	reg, err := regexp.Compile(args[0])
	if err != nil {
		log.Fatal(l.err, "\n", err)
	}
	for i, str := range data {
		if reg.MatchString(str) {
			fmt.Println(i + 1)
		}
	}
}
