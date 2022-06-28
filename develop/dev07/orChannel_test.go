package main

import (
	"fmt"
	"testing"
	"time"
)

func TestOrChannel(t *testing.T) {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
			fmt.Println("channel close")
		}()
		return c
	}

	start := time.Now()
	<-OrChannel(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(7*time.Second),
		sig(4*time.Second),
		sig(3*time.Second),
		sig(6*time.Second),
	)

	fmt.Printf("fone after %v", time.Since(start))

}
