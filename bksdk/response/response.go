package response

// /api/status
type Status []struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// /api/market/trades
type MarketTrades struct {
	Error  int                `json:"error"`
	Result MarketTradesResult `json:"result"`
}

type MarketTradesResult [][4]any

// /api/market/bids
type MarketBids struct {
	Error  int          `json:"error"`
	Result MarketResult `json:"result"`
}

// /api/market/asks
type MarketAsks struct {
	Error  int          `json:"error"`
	Result MarketResult `json:"result"`
}

type MarketResult [][5]any

// /api/market/books
type MarketBooks struct {
	Error  int               `json:"error"`
	Result MarketBooksResult `json:"result"`
}

type MarketBooksResult struct {
	Bids [][5]any `json:"bids"`
	Asks [][5]any `json:"asks"`
}

// /tradingview/history
type TradingviewHistory struct {
	C []float64 `json:"c"`
	H []float64 `json:"h"`
	L []float64 `json:"l"`
	O []float64 `json:"o"`
	S string    `json:"s"`
	T []int     `json:"t"`
	V []float64 `json:"v"`
}

// /api/market/depth
type MarketDepth struct {
	Asks [][]float64 `json:"asks"`
	Bids [][]float64 `json:"bids"`
}

// /api/market/ticker
type MarketTickerData struct {
	ID            int     `json:"id"`
	Last          float64 `json:"last"`
	LowestAsk     float64 `json:"lowestAsk"`
	HighestBid    float64 `json:"highestBid"`
	PercentChange float64 `json:"percentChange"`
	BaseVolume    float64 `json:"baseVolume"`
	QuoteVolume   float64 `json:"quoteVolume"`
	IsFrozen      int     `json:"isFrozen"`
	High24Hr      float64 `json:"high24hr"`
	Low24Hr       float64 `json:"low24hr"`
}

// /api/market/symbols
type MarketSymbols struct {
	Error  int                   `json:"error"`
	Result []MarketSymbolsResult `json:"result"`
}

type MarketSymbolsResult struct {
	ID     int    `json:"id"`
	Symbol string `json:"symbol"`
	Info   string `json:"info"`
}

// /api/v3/market/my-open-orders
type MyOpenOrder struct {
	Error  int                 `json:"error"`
	Result []MyOpenOrderResult `json:"result"`
}

type MyOpenOrderResult struct {
	ID       string  `json:"id"`
	Hash     string  `json:"hash"`
	Side     string  `json:"side"`
	Type     string  `json:"type"`
	Rate     int     `json:"rate"`
	Fee      float64 `json:"fee"`
	Credit   float64 `json:"credit"`
	Amount   float64 `json:"amount"`
	Receive  int     `json:"receive"`
	ParentID int     `json:"parent_id"`
	SuperID  int     `json:"super_id"`
	ClientID string  `json:"client_id"`
	Ts       int     `json:"ts"`
}

// /api/v3/market/my-order-history
type MyOrderHistory struct {
	Error      int                    `json:"error"`
	Result     []MyOrderHistoryResult `json:"result"`
	Pagination BKPaginate             `json:"pagination"`
}

type MyOrderHistoryResult struct {
	TxnID         string `json:"txn_id"`
	OrderID       string `json:"order_id"`
	Hash          string `json:"hash"`
	ParentOrderID string `json:"parent_order_id"`
	SuperOrderID  string `json:"super_order_id"`
	TakenByMe     bool   `json:"taken_by_me"`
	IsMaker       bool   `json:"is_maker"`
	Side          string `json:"side"`
	Type          string `json:"type"`
	Rate          string `json:"rate"`
	Fee           string `json:"fee"`
	Credit        string `json:"credit"`
	Amount        string `json:"amount"`
	Ts            int    `json:"ts"`
}

type BKPaginate struct {
	Page int `json:"page"`
	Last int `json:"last"`
	Next int `json:"next"`
	Prev int `json:"prev"`
}

// /api/v3/market/order-info
type OrderInfo struct {
	Error  int             `json:"error"`
	Result OrderInfoResult `json:"result"`
}

type OrderInfoResult struct {
	ID            string                   `json:"id"`
	First         string                   `json:"first"`
	Parent        string                   `json:"parent"`
	Last          string                   `json:"last"`
	Amount        string                   `json:"amount"`
	Rate          int                      `json:"rate"`
	Fee           int                      `json:"fee"`
	Credit        int                      `json:"credit"`
	Filled        float64                  `json:"filled"`
	Total         int                      `json:"total"`
	Status        string                   `json:"status"`
	PartialFilled bool                     `json:"partial_filled"`
	Remaining     int                      `json:"remaining"`
	History       []OrderInfoResultHistory `json:"history"`
}

type OrderInfoResultHistory struct {
	Amount    float64 `json:"amount"`
	Credit    float64 `json:"credit"`
	Fee       float64 `json:"fee"`
	Hash      string  `json:"hash"`
	ID        string  `json:"id"`
	Rate      int     `json:"rate"`
	Timestamp int64   `json:"timestamp"`
	TxnID     string  `json:"txn_id"`
}

type TradingCredit struct {
	Error  int     `json:"error"`
	Result float64 `json:"result"`
}

type Limits struct {
	Error  int          `json:"error"`
	Result LimitsResult `json:"result"`
}

type LimitsResult struct {
	Limits struct {
		Crypto struct {
			Deposit  float64 `json:"deposit"`
			Withdraw float64 `json:"withdraw"`
		} `json:"crypto"`
		Fiat struct {
			Deposit  float32 `json:"deposit"`
			Withdraw float32 `json:"withdraw"`
		} `json:"fiat"`
	} `json:"limits"`
	Usage struct {
		Crypto struct {
			Deposit               float64 `json:"deposit"`
			Withdraw              float64 `json:"withdraw"`
			DepositPercentage     float64 `json:"deposit_percentage"`
			WithdrawPercentage    float64 `json:"withdraw_percentage"`
			DepositThbEquivalent  float64 `json:"deposit_thb_equivalent"`
			WithdrawThbEquivalent float64 `json:"withdraw_thb_equivalent"`
		} `json:"crypto"`
		Fiat struct {
			Deposit            float32 `json:"deposit"`
			Withdraw           float32 `json:"withdraw"`
			DepositPercentage  float32 `json:"deposit_percentage"`
			WithdrawPercentage float32 `json:"withdraw_percentage"`
		} `json:"fiat"`
	} `json:"usage"`
	Rate float32 `json:"rate"`
}

type Wallet struct {
	Error  int          `json:"error"`
	Result WalletResult `json:"result"`
}
type WalletResult struct {
	Thb float64 `json:"THB"`
	Btc float64 `json:"BTC"`
	Eth float64 `json:"ETH"`
}

type Balances struct {
	Error  int           `json:"error"`
	Result BalanceResult `json:"result"`
}

type BalanceResult map[string]BalanceMapResult

type BalanceMapResult struct {
	Available float64 `json:"available"`
	Reserved  float64 `json:"reserved"`
}

type WsToken struct {
	Error  int    `json:"error"`
	Result string `json:"result"`
}

type InternalWithdraw struct {
	Error  int                    `json:"error"`
	Result InternalWithdrawResult `json:"result"`
}

type InternalWithdrawResult struct {
	Txn string  `json:"txn"`
	Adr string  `json:"adr"`
	Mem string  `json:"mem"`
	Cur string  `json:"cur"`
	Amt float64 `json:"amt"`
	Fee float64 `json:"fee"`
	Ts  int     `json:"ts"`
}

type DepositHistory struct {
	Error      int                    `json:"error"`
	Result     []DepositHistoryResult `json:"result"`
	Pagination BKPaginate             `json:"pagination"`
}

type DepositHistoryResult struct {
	Hash          string  `json:"hash"`
	Currency      string  `json:"currency"`
	Amount        float64 `json:"amount"`
	FromAddress   string  `json:"from_address"`
	ToAddress     string  `json:"to_address"`
	Confirmations int     `json:"confirmations"`
	Status        string  `json:"status"`
	Time          int     `json:"time"`
}

type WithdrawHistory struct {
	Error      int                     `json:"error"`
	Result     []WithdrawHistoryResult `json:"result"`
	Pagination BKPaginate              `json:"pagination"`
}

type WithdrawHistoryResult struct {
	TxnID    string  `json:"txn_id"`
	Hash     string  `json:"hash"`
	Currency string  `json:"currency"`
	Amount   string  `json:"amount"`
	Fee      float64 `json:"fee"`
	Address  string  `json:"address"`
	Status   string  `json:"status"`
	Time     int     `json:"time"`
}

type PlaceBid struct {
	Error  int            `json:"error"`
	Result PlaceBidResult `json:"result"`
}

type PlaceBidResult struct {
	ID   string  `json:"id"`
	Hash string  `json:"hash"`
	Typ  string  `json:"typ"`
	Amt  int     `json:"amt"`
	Rat  int     `json:"rat"`
	Fee  float64 `json:"fee"`
	Cre  float64 `json:"cre"`
	Rec  float64 `json:"rec"`
	Ts   int     `json:"ts"`
	Ci   string  `json:"ci"`
}

type PlaceAsk struct {
	Error  int            `json:"error"`
	Result PlaceAskResult `json:"result"`
}

type PlaceAskResult struct {
	ID   string  `json:"id"`
	Hash string  `json:"hash"`
	Typ  string  `json:"typ"`
	Amt  float64 `json:"amt"`
	Rat  int     `json:"rat"`
	Fee  float64 `json:"fee"`
	Cre  float64 `json:"cre"`
	Rec  int     `json:"rec"`
	Ts   int     `json:"ts"`
	Ci   string  `json:"ci"`
}

type CancelOrder struct {
	Error int `json:"error"`
}

type CryptoAddresses struct {
	Error      int                     `json:"error"`
	Result     []CryptoAddressesResult `json:"result"`
	Pagination BKPaginate              `json:"pagination"`
}

type CryptoAddressesResult struct {
	Currency string `json:"currency"`
	Address  string `json:"address"`
	Tag      int    `json:"tag"`
	Time     int    `json:"time"`
}

type CryptoGenerateAddress struct {
	Error  int                           `json:"error"`
	Result []CryptoGenerateAddressResult `json:"result"`
}

type CryptoGenerateAddressResult struct {
	Currency string `json:"currency"`
	Address  string `json:"address"`
	Memo     string `json:"memo"`
}

type CryptoWithdraw struct {
	Error  int                  `json:"error"`
	Result CryptoWithdrawResult `json:"result"`
}

type CryptoWithdrawResult struct {
	Txn string  `json:"txn"`
	Adr string  `json:"adr"`
	Mem string  `json:"mem"`
	Cur string  `json:"cur"`
	Amt float64 `json:"amt"`
	Fee float64 `json:"fee"`
	Ts  int     `json:"ts"`
}

type FiatAccounts struct {
	Error      int                  `json:"error"`
	Result     []FiatAccountsResult `json:"result"`
	Pagination BKPaginate           `json:"pagination"`
}

type FiatAccountsResult struct {
	ID   string `json:"id"`
	Bank string `json:"bank"`
	Name string `json:"name"`
	Time int    `json:"time"`
}

type FiatWithdraw struct {
	Error  int                `json:"error"`
	Result FiatWithdrawResult `json:"result"`
}

type FiatWithdrawResult struct {
	Txn string `json:"txn"`
	Acc string `json:"acc"`
	Cur string `json:"cur"`
	Amt int    `json:"amt"`
	Fee int    `json:"fee"`
	Rec int    `json:"rec"`
	Ts  int    `json:"ts"`
}

type FiatDepositHistory struct {
	Error      int                        `json:"error"`
	Result     []FiatDepositHistoryResult `json:"result"`
	Pagination BKPaginate                 `json:"pagination"`
}

type FiatDepositHistoryResult struct {
	TxnID    string  `json:"txn_id"`
	Currency string  `json:"currency"`
	Amount   float64 `json:"amount"`
	Status   string  `json:"status"`
	Time     int     `json:"time"`
}

type FiatWithdrawHistory struct {
	Error      int                         `json:"error"`
	Result     []FiatWithdrawHistoryResult `json:"result"`
	Pagination BKPaginate                  `json:"pagination"`
}

type FiatWithdrawHistoryResult struct {
	TxnID    string `json:"txn_id"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	Fee      int    `json:"fee"`
	Status   string `json:"status"`
	Time     int    `json:"time"`
}

// websocket response
type WsTrade struct {
	Amt    float64 `json:"amt"`
	Bid    string  `json:"bid"`
	Rat    float64 `json:"rat"`
	Sid    string  `json:"sid"`
	Stream string  `json:"stream"`
	Sym    string  `json:"sym"`
	Ts     int     `json:"ts"`
	Txn    string  `json:"txn"`
}

type WsTicker struct {
	Stream         string  `json:"stream"`
	ID             int     `json:"id"`
	Last           float64 `json:"last"`
	LowestAsk      float64 `json:"lowestAsk"`
	LowestAskSize  float64 `json:"lowestAskSize"`
	HighestBid     float64 `json:"highestBid"`
	HighestBidSize float64 `json:"highestBidSize"`
	Change         float64 `json:"change"`
	PercentChange  float64 `json:"percentChange"`
	BaseVolume     float64 `json:"baseVolume"`
	QuoteVolume    float64 `json:"quoteVolume"`
	IsFrozen       int     `json:"isFrozen"`
	High24Hr       float64 `json:"high24hr"`
	Low24Hr        float64 `json:"low24hr"`
	Open           float64 `json:"open"`
	Close          float64 `json:"close"`
}
