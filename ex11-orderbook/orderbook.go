package orderbook

import "sort"

type Orderbook struct {
	Bids []*Order
	Asks []*Order
}

func New() *Orderbook {
	ob := &Orderbook{}
	ob.Bids = []*Order{}
	ob.Asks = []*Order{}
	return ob
}

func (ob *Orderbook) AddBid(order *Order) {
	index := sort.Search(len(ob.Bids), func(i int) bool { return ob.Bids[i].Price > order.Price })
	ob.Bids = append(ob.Bids, order)
	copy(ob.Bids[index+1:], ob.Bids[index:])
	ob.Bids[index] = order
}

func (ob *Orderbook) AddAsk(order *Order) {
	index := sort.Search(len(ob.Asks), func(i int) bool { return ob.Asks[i].Price < order.Price })
	ob.Asks = append(ob.Asks, order)
	copy(ob.Asks[index+1:], ob.Asks[index:])
	ob.Asks[index] = order
}

func (ob *Orderbook) TradeAsks(order *Order) ([]*Trade, *Order) {
	trades := []*Trade{}

	for i := 0; i < len(ob.Asks); i++ {
		a := ob.Asks[i]
		if order.Price <= a.Price {
			if a.Volume <= order.Volume {
				trades = append(trades, &Trade{Bid: order, Ask: a, Volume: a.Volume, Price: a.Price})
				order.Volume -= a.Volume
				a.Volume = 0
				ob.Asks = append(ob.Asks[:i], ob.Asks[i+1:]...)
				i--
			} else {
				trades = append(trades, &Trade{Bid: order, Ask: a, Volume: order.Volume, Price: a.Price})
				a.Volume -= order.Volume
				order.Volume = 0
				return trades, nil
			}
		} else {
			break
		}
	}

	if order.Kind == 1 {
		return trades, order
	} else {
		ob.AddBid(order)
	}
	return trades, nil
}

func (ob *Orderbook) TradeBids(order *Order) ([]*Trade, *Order) {
	trades := []*Trade{}

	for i := 0; i < len(ob.Bids); i++ {
		bid := ob.Bids[i]
		if order.Price == 0 || bid.Price <= order.Price {
			if bid.Volume <= order.Volume {
				trades = append(trades, &Trade{Bid: bid, Ask: order, Volume: bid.Volume, Price: bid.Price})
				order.Volume -= bid.Volume
				bid.Volume = 0
				ob.Bids = append(ob.Bids[:i], ob.Bids[i+1:]...)
				i -= 1
			} else {
				trades = append(trades, &Trade{Bid: bid, Ask: order, Volume: order.Volume, Price: bid.Price})
				bid.Volume -= order.Volume
				order.Volume = 0
			}
			if order.Volume == 0 {
				return trades, nil
			}
		} else {
			break
		}
	}

	if order.Kind == 1 {
		return trades, order
	} else {
		ob.AddAsk(order)
	}
	return trades, nil
}

func (ob *Orderbook) Match(order *Order) ([]*Trade, *Order) {
	switch order.Side {
	case SideAsk:
		return ob.TradeAsks(order)
	case SideBid:
		return ob.TradeBids(order)
	}

	return nil, nil
}
