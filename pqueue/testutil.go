package pqueue

import (
	"github.com/fmstephe/matching_engine/trade"
	"testing"
)

var orderMaker = trade.NewOrderMaker()

func AllSameBuy(t *testing.T, h Q, verifyQ func(*testing.T, Q)) {
	if h.BuySell() != trade.BUY {
		t.Errorf("Expecting BUY queue")
	}
	for i := 20; i > 0; i-- {
		h.Push(orderMaker.MkPricedBuy(1))
	}
	verifyQ(t, h)
	for i := 1; h.Size() > 0; i++ {
		x := h.Pop()
		verifyQ(t, h)
		if x.Price != 1 {
			t.Errorf("%d.th Pop got %d; want %d", i, x, 0)
		}
	}
}

func AllSameSell(t *testing.T, h Q, verifyQ func(*testing.T, Q)) {
	if h.BuySell() != trade.SELL {
		t.Errorf("Expecting SELL queue")
	}
	for i := 20; i > 0; i-- {
		h.Push(orderMaker.MkPricedSell(1))
	}
	verifyQ(t, h)
	for i := 1; h.Size() > 0; i++ {
		x := h.Pop()
		verifyQ(t, h)
		if x.Price != 1 {
			t.Errorf("%d.th Pop got %d; want %d", i, x, 0)
		}
	}
}

func DescendingBuy(t *testing.T, h Q, verifyQ func(*testing.T, Q)) {
	for i := int32(20); i > 0; i-- {
		h.Push(orderMaker.MkPricedBuy(i))
	}
	verifyQ(t, h)
	for i := int32(20); h.Size() > 0; i-- {
		x := h.Pop()
		verifyQ(t, h)
		if x.Price != i {
			t.Errorf("%d.th Pop got %d; want %d", i, x, i)
		}
	}
}

func DescendingSell(t *testing.T, h Q, verifyQ func(*testing.T, Q)) {
	for i := int32(20); i > 0; i-- {
		h.Push(orderMaker.MkPricedSell(i))
	}
	verifyQ(t, h)
	for i := int32(1); h.Size() > 0; i++ {
		x := h.Pop()
		verifyQ(t, h)
		if x.Price != i {
			t.Errorf("%d.th Pop got %d; want %d", i, x, i)
		}
	}
}

func AscendingBuy(t *testing.T, h Q, verifyQ func(*testing.T, Q)) {
	for i := int32(1); i <= 20; i++ {
		h.Push(orderMaker.MkPricedBuy(i))
	}
	verifyQ(t, h)
	for i := int32(20); h.Size() > 0; i-- {
		x := h.Pop()
		verifyQ(t, h)
		if x.Price != i {
			t.Errorf("%d.th Pop got %d; want %d", i, x, i)
		}
	}
}

func AscendingSell(t *testing.T, h Q, verifyQ func(*testing.T, Q)) {
	for i := int32(1); i <= 20; i++ {
		h.Push(orderMaker.MkPricedSell(i))
	}
	verifyQ(t, h)
	for i := int32(1); h.Size() > 0; i++ {
		x := h.Pop()
		verifyQ(t, h)
		if x.Price != i {
			t.Errorf("%d.th Pop got %d; want %d", i, x, i)
		}
	}
}

func BuyRandomPushPop(t *testing.T, h Q, verifyQ func(*testing.T, Q)) {
	size := 10000
	priceRange := int32(500)
	priceBase := int32(1000)
	buys := make([]*trade.Order, 0, size)
	for i := 0; i < size; i++ {
		b := orderMaker.MkPricedBuy(orderMaker.Rand32(priceRange) + priceBase)
		buys = append(buys, b)
		h.Push(b)
		verifyQ(t, h)
	}
	leastPrice := priceRange + priceBase + 1
	for i := 0; i < size; i++ {
		b := h.Pop()
		if b.Price > leastPrice {
			t.Errorf("Buy Pop reveals out of order buy order")
		}
		leastPrice = b.Price
		verifyQ(t, h)
	}
}

func SellRandomPushPop(t *testing.T, h Q, verifyQ func(*testing.T, Q)) {
	size := 10000
	priceRange := int32(500)
	priceBase := int32(1000)
	buys := make([]*trade.Order, 0, size)
	for i := 0; i < size; i++ {
		b := orderMaker.MkPricedSell(orderMaker.Rand32(priceRange) + priceBase)
		buys = append(buys, b)
		h.Push(b)
		verifyQ(t, h)
	}
	greatestPrice := int32(0)
	for i := 0; i < size; i++ {
		s := h.Pop()
		if s.Price < greatestPrice {
			t.Errorf("Sell Pop reveals out of order sell order")
		}
		greatestPrice = s.Price
		verifyQ(t, h)
	}
}