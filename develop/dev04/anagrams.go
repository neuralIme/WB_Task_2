package main

import (
	"sort"
	"strings"
)

// Anagrams с помощью вложенных циклов конвертирует полученные слова в мапу анаграмм
func Anagrams(inp []string) map[string][]string {
	res := make(map[string][]string)

	for k, v := range inp { // внешний цикл

		lov := strings.ToLower(v) // приведение к нижнему регистру
		key1 := conv([]rune(lov)) // создание первого ключа для сравнения
		res[lov] = []string{}     // фиксация слова в мапе

		for i := k; i < len(inp); i++ { // внутренний цикл

			low := strings.ToLower(inp[i])  // приведение к нижнему регистру
			key2 := conv([]rune(low))       // создание второго ключа для сравнения
			if key1 == key2 && lov != low { // если ключи подходят, а слова различаются добавление в мапу
				res[lov] = append(res[lov], low)
				inp[i] = ""
			}
		}
		if len(res[v]) == 0 {
			delete(res, v) // удаление слов-одиночек
		}
	}

	for _, v := range res {
		sort.Strings(v) // сортировка массивов по возрастанию
	}

	return res
}

// conv сортирует руны в слове для ключа
func conv(ru []rune) string {
	for i := range ru {
		for j := i; j > 0 && ru[j] > ru[j-1]; j-- {
			ru[j], ru[j-1] = ru[j-1], ru[j]
		}
	}
	return string(ru)
}
