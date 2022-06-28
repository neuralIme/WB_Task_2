package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("I am app1")
		time.Sleep(time.Second * 4)
	}

}
