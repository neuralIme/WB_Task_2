package main

import (
	"fmt"
	"testing"
)

func TestAnagrams(t *testing.T) {
	data := []string{
		"корабль", "МАТерик", "пятак",
		"ЛИСток", "пятка", "КЕРамит",
		"слиток", "ТЯПка", "столик",
		"метрика", "термика", ""}

	fmt.Println(Anagrams(data))
}
