package main

import (
	"./barbershop"
	"fmt"
	"time"
)

func main() {
	barber := barbershop.NewBarber("Alexander")
	shop := barbershop.NewShop(barber, 4)
	go barber.ManageShop(shop)
	time.Sleep(100 * time.Millisecond)

	clients := []string{"Steve", "Taras", "Yulia", "Ihor", "Denis", "Stas"}
	for _, c := range clients {
		client := barbershop.NewClient(c)
		go client.EnterShop(shop)
	}

	fmt.Scanln()
}
