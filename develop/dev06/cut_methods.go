package main

import (
	"fmt"
	"strconv"
	"strings"
)

// выбрать поля (колонки)
type fields struct {
	err string
}

// redact выводит заданные поля с разделителем по умолчанию
func (f *fields) redact(cmd []string) {
	f.err = "некорректный ввод"
	if len(cmd) < 1 {
		for i := range data {
			fmt.Println(strings.Join(data[i], "\t"))
		}
	} else if len(cmd) == 1 {
		if len(cmd[0]) == 1 {
			n, err := strconv.Atoi(cmd[0])
			if err != nil {
				println(f.err)
				return
			}
			for i := range data {
				if n > len(data[i]) {
					continue
				}
				fmt.Println(data[i][n-1])
			}
		} else {
			switch cmd[0][1] {
			case ',':
				s := strings.Split(cmd[0], ",")
				for i := range data {
					for _, j := range s {
						num, err := strconv.Atoi(j)
						if err != nil {
							println(f.err)
							return
						}
						if num > len(data[i]) {
							continue
						}
						if num == len(data[i]) {
							fmt.Print(data[i][num-1])
						} else {
							fmt.Print(data[i][num-1], "\t")
						}
					}
					fmt.Print("\n")
				}
			case '-':
				a, err := strconv.Atoi(string(cmd[0][0]))
				if err != nil {
					println(f.err)
					return
				}
				b, err := strconv.Atoi(string(cmd[0][2]))
				if err != nil {
					println(f.err)
					return
				}
				for i := range data {
					for j := a; j <= b; j++ {
						if j > len(data[i]) {
							continue
						}
						if j == len(data[i]) {
							fmt.Print(data[i][j-1])
						} else {
							fmt.Print(data[i][j-1], "\t")
						}
					}
					fmt.Print("\n")
				}
			default:
				println(f.err)
				return
			}

		}
	} else {
		println(f.err)
		return
	}
}

// использовать другой разделитель
type delimiter struct {
	err string
}

// redact выводит указанные поля с заданным разделителем
func (d *delimiter) redact(cmd []string) {
	f.err = "некорректный ввод"
	if len(cmd) < 3 {
		if len(cmd) == 0 {
			println(f.err)
			return
		}
		for i := range data {
			fmt.Println(strings.Join(data[i], cmd[0]))
		}
	} else if len(cmd) == 3 {
		if len(cmd[2]) == 1 {
			n, err := strconv.Atoi(cmd[0])
			if err != nil {
				println(f.err)
				return
			}
			for i := range data {
				if n > len(data[i]) {
					continue
				}
				fmt.Println(data[i][n-1])
			}
		} else {
			switch cmd[2][1] {
			case ',':
				s := strings.Split(cmd[2], ",")
				for i := range data {
					for _, j := range s {
						num, err := strconv.Atoi(j)
						if err != nil {
							println(f.err)
							return
						}
						if num > len(data[i]) {
							continue
						}
						if num == len(data[i]) {
							fmt.Print(data[i][num-1])
						} else {
							fmt.Print(data[i][num-1], cmd[0])
						}
					}
					fmt.Print("\n")
				}
			case '-':
				a, err := strconv.Atoi(string(cmd[2][0]))
				if err != nil {
					println(f.err)
					return
				}
				b, err := strconv.Atoi(string(cmd[2][2]))
				if err != nil {
					println(f.err)
					return
				}
				for i := range data {
					for j := a; j <= b; j++ {
						if j > len(data[i]) {
							continue
						}
						if j == len(data[i]) {
							fmt.Print(data[i][j-1])
						} else {
							fmt.Print(data[i][j-1], cmd[0])
						}
					}
					fmt.Print("\n")
				}
			default:
				println(f.err)
				return
			}

		}
	} else {
		println(f.err)
		return
	}
}

// только строки с разделителем
type separated struct {
	err string
}

// redact выводит строки имеющие разделитель
func (s *separated) redact(cmd []string) {
	s.err = "некорректный ввод"
	if len(cmd) < 1 {
		for i := range data {
			if len(data[i]) == 1 {
				continue
			}
			fmt.Println(strings.Join(data[i], "\t"))
		}
	} else if len(cmd) == 1 {
		if len(cmd[0]) == 1 {
			n, err := strconv.Atoi(cmd[0])
			if err != nil {
				println(f.err)
				return
			}
			for i := range data {
				if len(data[i]) == 1 {
					continue
				}
				if n > len(data[i]) {
					continue
				}
				fmt.Println(data[i][n-1])
			}
		} else {
			switch cmd[0][1] {
			case ',':
				s := strings.Split(cmd[0], ",")
				for i := range data {
					for _, j := range s {
						if len(data[i]) == 1 {
							continue
						}
						num, err := strconv.Atoi(j)
						if err != nil {
							println(f.err)
							return
						}
						if num > len(data[i]) {
							continue
						}
						if num == len(data[i]) {
							fmt.Print(data[i][num-1])
						} else {
							fmt.Print(data[i][num-1], "\t")
						}
					}
					fmt.Print("\n")
				}
			case '-':
				a, err := strconv.Atoi(string(cmd[0][0]))
				if err != nil {
					println(f.err)
					return
				}
				b, err := strconv.Atoi(string(cmd[0][2]))
				if err != nil {
					println(f.err)
					return
				}
				for i := range data {
					for j := a; j <= b; j++ {
						if len(data[i]) == 1 {
							continue
						}
						if j > len(data[i]) {
							continue
						}
						if j == len(data[i]) {
							fmt.Print(data[i][j-1])
						} else {
							fmt.Print(data[i][j-1], "\t")
						}
					}
					fmt.Print("\n")
				}
			default:
				println(f.err)
				return
			}

		}
	} else {
		println(f.err)
		return
	}
}
