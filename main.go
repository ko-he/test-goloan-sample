package main

import (
	"fmt"

	"github.com/ko-he/test-goloan-sample/infrastructure/persistence/csvstore"
	"github.com/ko-he/test-goloan-sample/service"
)

func main() {
	cr := csvstore.NewClientRepository()
	cs := service.NewClientService(cr)
	clients, _ := cs.SearchByPrefacture("京都")

	for _, c := range clients {
		fmt.Println(c.ID, c.Name, c.Address)
	}
}
