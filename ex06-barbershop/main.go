package main

import (
	"fmt"

	"./barbershop"
)

func main() {
	barber := barbershop.NewBarber("Alexander")
	shop := barbershop.NewShop(barber, 4)
	go barber.ManageShop(shop)

	clients := []string{"Steve", "Taras", "Yulia", "Ihor", "Denis", "Stas"}
	for _, c := range clients {
		client := barbershop.NewClient(c)
		go client.EnterShop(shop)
	}

	fmt.Scanln()
}
