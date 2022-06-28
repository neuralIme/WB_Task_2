package main

import (
	"fmt"
	"log"
	"testing"
)

func TestUnpacking(t *testing.T) {
	teststring := []string{
		"a4bc2d5e",
		"abcd",
		"45",
		"",
		"qwe\\4\\5",
		"qwe\\45",
		"qwe\\\\5",
		"aaa0b",
	}

	for _, i := range teststring {
		res, err := Unpacking(i)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(res)
	}

}
