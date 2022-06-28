package main

import (
	"errors"
	"fmt"
	"unicode"
)

var err = errors.New("некорректная строка")

// Unpacking распаковывает полученную строку, итерируясь по каждому символу
func Unpacking(str string) (string, error) {
	ru := []rune(str)
	var by []byte
	var n int
	var backslash bool

	for i, item := range ru {
		if unicode.IsDigit(item) && i == 0 {
			return "", err
		}
		if unicode.IsDigit(item) && unicode.IsDigit(ru[i-1]) && ru[i-2] != '\\' {
			return "", err
		}
		if item == '\\' && !backslash {
			backslash = true
			continue
		}
		if backslash && unicode.IsLetter(item) {
			return "", err
		}
		if backslash {
			by = append(by, byte(item))
			backslash = false
			continue
		}
		if unicode.IsDigit(item) {
			n = int(item - '0')
			fmt.Println(n)
			if n == 0 {
				by = by[:len(by)-1]
				continue
			}
			for j := 0; j < n-1; j++ {
				by = append(by, byte(ru[i-1]))
			}
			continue
		}
		by = append(by, byte(item))
	}

	return string(by), nil
}
