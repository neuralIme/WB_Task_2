package main

import (
	"github.com/beevik/ntp"
	"log"
	"time"
)

// GetTimeNow возвращает текущее время или выводит ошибку в терминал
func GetTimeNow() time.Time {

	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		log.Fatal(err.Error())
	}

	return time
}
