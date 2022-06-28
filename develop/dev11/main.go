package main

import (
	"calendar/api"
	"calendar/domain"
)

func main() {

	storage := domain.NewStorage()

	h := &api.Handler{Storage: storage}
	h.StartServer()

}
