package barbershop

import (
	"fmt"
	"time"
)

func (b *Barber) ManageShop(shop *Shop) {
	for {
		// manage a shop with a waiting room
		select {
		case c := <-shop.WaitingRoom:
			fmt.Printf("%s is cutting %s`s hair\n", b.GetName(), c.GetName())
			time.Sleep(CUTTING_TIME * time.Millisecond)
			fmt.Printf("%s finished work with %s\n", b.GetName(), c.GetName())
		default:
			fmt.Printf("%s is sleeping...\n", b.GetName())
			c := <-b.WakeUp
			fmt.Printf("%s waked by %s\n", b.GetName(), c.GetName())
		}
	}
}

func (c *Client) EnterShop(shop *Shop) {
	for i := 0; i < FIND_SEAT_TRIES; i++ {
		// checking for returning client
		if i > 0 {
			fmt.Printf("%s returned\n", c.GetName())
		}
		// checking for free seat
		select {
		case shop.WaitingRoom <- c:
			fmt.Printf("%s found a seat\n", c.GetName())
			// wake up barber
			select {
			case shop.Barber.WakeUp <- c:
			default:
			}
			return
		default:
			fmt.Printf("%s will back later...\n", c.GetName())
			time.Sleep(GO_OUT_TIME * time.Millisecond)
		}
	}
	fmt.Printf("%s didn`t found any seat.\n", c.GetName())
}
