package barbershop

const (
	CUTTING_TIME    = 10
	FIND_SEAT_TRIES = 3
	GO_OUT_TIME     = 20
)

type Barber struct {
	name   string
	WakeUp chan *Client
}

type Client struct {
	name string
}

type Shop struct {
	Barber      *Barber
	WaitingRoom chan *Client
}

func NewBarber(name string) *Barber {
	b := &Barber{}
	b.name = name
	b.WakeUp = make(chan *Client)
	return b
}

func NewClient(name string) *Client {
	c := &Client{}
	c.name = name
	return c
}

func NewShop(barber *Barber, seats int) *Shop {
	shop := &Shop{}
	shop.Barber = barber
	shop.WaitingRoom = make(chan *Client, seats)
	return shop
}

func (b *Barber) GetName() string {
	return b.name
}

func (c *Client) GetName() string {
	return c.name
}
