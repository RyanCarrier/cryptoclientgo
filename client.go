package cryptoclientgo

//CryptoClient is the generic crypto currency client
type CryptoClient struct{ AbstractClient }

//New gets a new cryptoclient
func New(client AbstractClient) CryptoClient {
	return CryptoClient{client}
}

//AbstractClient is an abstract client for basic cryptoClient methods
type AbstractClient interface {
	//Public
	GetPrimaryCurrencies() ([]string, error)
	GetSecondaryCurrencies() ([]string, error)
	Tick(Currency CurrencyPair) (Tick, error)
	GetOrderBook(Currency CurrencyPair) (OrderBook, error)
	GetRecentTrades(Currency CurrencyPair, historyAmount int) (RecentTrades, error)
	//ExpectedMarketValueBuy(PrimaryCurrency, SecondaryCurrency string, amountOfToCurrency int64) (int64, error)
	//ExpectedMarketValueSell(PrimaryCurrency, SecondaryCurrency string, amountOfFromCurrency int64) (int64, error)
	//Private
	PlaceLimitBuyOrder(Currency CurrencyPair, amount int64, price int64) (PlacedOrder, error)
	PlaceMarketBuyOrder(Currency CurrencyPair, amount int64) (PlacedOrder, error)
	PlaceLimitSellOrder(Currency CurrencyPair, amount int64, price int64) (PlacedOrder, error)
	PlaceMarketSellOrder(Currency CurrencyPair, amount int64) (PlacedOrder, error)
	CancelOrder(OrderID int) error
	GetOrderDetails(OrderID int) (OrderDetails, error)
	GetOpenOrders(CurrencyPair) (OrdersDetails, error)
	GetBalance(Currency string) (AccountBalance, error)
	GetBalances() (AccountBalances, error)
	GetPrimaryCurrencyDepositAddress(Currency string) (CurrencyAddress, error)
	WithdrawCurrency(Currency, to string, amount int64) error
	//Custom
	GetTransactionCost(Currency CurrencyPair) (Cost, error)
	GetWithdrawCost(Currency string) (Cost, error)
	GetDepositCost(Currency string) (Cost, error)
}
