package main

import (
	"sync"
)

// OrChannel сливает полученные каналы в один и возвращает его;
// передает в него данные даже при закрытии одного из составляющих каналов
func OrChannel(channels ...<-chan interface{}) <-chan interface{} {
	wg := new(sync.WaitGroup)
	out := make(chan interface{})

	wg.Add(len(channels))
	merge := func(ch <-chan interface{}) {
		defer wg.Done()
		for i := range ch {
			out <- i
		}
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for _, ch := range channels {
		go merge(ch)
	}

	return out
}
