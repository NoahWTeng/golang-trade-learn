package main

import "fmt"

const (
	BITFINEX_API_URL          = "https://api.bitfinex.com/v1/"
	BITFINEX_TICKER           = "pubticker/"
	BITFINEX_STATS            = "stats/"
	BITFINEX_ORDERBOOK        = "book/"
	BITFINEX_TRADES           = "trades/"
	BITFINEX_SYMBOLS          = "symbols/"
	BITFINEX_SYMBOLS_DETAILS  = "symbols_details/"
	BITFINEX_DEPOSIT          = "deposit/new"
	BITFINEX_ORDER_NEW        = "order/new"
	BITFINEX_ORDER_CANCEL     = "order/cancel"
	BITFINEX_ORDER_CANCEL_ALL = "order/cancel/all"
	BITFINEX_ORDER_STATUS     = "order/status"
	BITFINEX_ORDERS           = "orders"
	BITFINEX_POSITIONS        = "positions"
	BITFINEX_CLAIM_POSITION   = "position/claim"
	BITFINEX_HISTORY          = "history"
	BITFINEX_TRADE_HISTORY    = "mytrades"
)

type BitfinexStats struct {
	Period int64
	Volume string
}

type BitfinexTicker struct {
	Mid, Bid, Ask, Last_price, Low, High, Volume, Timestamp string
}

type BookStructure struct {
	Price, Amount, Timestamp string
}

type BitfinexOrderbook struct {
	Bids []BookStructure
	Asks []BookStructure
}

type TradeStructure struct {
	Timestamp, Tid                int64
	Price, Amount, Exchange, Type string
}

type SymbolsDetails struct {
	Pair, Initial_margin, Minimum_margin, Maximum_order_size, Minimum_order_size, Expiration string
	Price_precision                                                                          int
}

type Bitfinex struct {
	APIKey, APISecret string
	Ticker            BitfinexTicker
	Stats             []BitfinexStats
	Orderbook         BitfinexOrderbook
	Trades            []TradeStructure
	SymbolsDetails    []SymbolsDetails
}

func (b *Bitfinex) GetTicker(symbol string) bool {

	err := SendHttpRequest(BITFINEX_API_URL+BITFINEX_TICKER+symbol, true, &b.Ticker)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
