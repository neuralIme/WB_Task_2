package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("I am app2")
		time.Sleep(time.Second * 4)
	}
}
