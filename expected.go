package cryptoclientgo

import (
	"errors"
	"math/big"
)

var ErrMath = errors.New("MATH ERROR, most likely BigInt too large for int64")

//ExpectedMarketValueBuy gets the expected value cost from PrimaryCurrency to achieve the specified amount of SecondaryCurrency
func (c CryptoClient) ExpectedMarketValueBuy(Currency CurrencyPair, amountOfSecondary int64) (int64, error) {
	return c.expectedMarketValue(Currency, amountOfSecondary, true)
}

//ExpectedMarketValueSell gets the expected value recieved of PrimaryCurrency by selling the specified amount of SecondaryCurrency
func (c CryptoClient) ExpectedMarketValueSell(Currency CurrencyPair, amountOfSecondary int64) (int64, error) {
	return c.expectedMarketValue(Currency, amountOfSecondary, false)
}

func (c CryptoClient) expectedMarketValue(Currency CurrencyPair, amt int64, buy bool) (int64, error) {
	order, err := c.GetOrderBook(Currency)
	if err != nil {
		return 0, errors.New("Failed to get open orders;" + err.Error())
	}
	if buy {
		//if we are buying we want best sell orders
		return order.SellOrders.getBestBuy(amt)
	}
	return order.BuyOrders.getBestSell(amt)
}

func (o Orders) getBestBuy(amt int64) (int64, error) {
	o.SortBuy()
	return o.getBest(amt)
}

func (o Orders) getBestSell(amt int64) (int64, error) {
	o.SortSell()
	return o.getBest(amt)
}

func (o Orders) getBest(amt int64) (int64, error) {
	//spew.Dump(o)
	BIGamt := big.NewInt(amt)
	BIGtotal := big.NewInt(0)
	temp := big.NewInt(0)
	for _, order := range o {
		BIGVolume := big.NewInt(order.Volume)
		BIGPrice := big.NewInt(order.Price)
		// fmt.Printf("Volume\t\tPrice\t\tamt\t\ttotal\n%+v\t%+v\t%+v\t%+v\n\n", order.Volume, order.Price, BIGamt, BIGtotal)
		// fmt.Println(BIGVolume, BIGamt)
		if (BIGVolume).Cmp(BIGamt) >= 0 { //-1 if BIGVolume smaller 0 or +1 if equal or greater
			BIGtotal.Add(BIGtotal, temp.Div(temp.Mul(BIGamt, BIGPrice), bigMultiplier))
			if BIGtotal.IsInt64() {
				return BIGtotal.Int64(), nil
			}
			return 0, ErrMath
		}
		BIGamt.Sub(BIGamt, BIGVolume)
		BIGtotal.Add(BIGtotal, temp.Div(temp.Mul(BIGVolume, BIGPrice), bigMultiplier))
	}
	return 0, errors.New("Not enough volume in open orders")
}
