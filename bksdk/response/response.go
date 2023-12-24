package response

// /api/status
type Status []struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

// /api/market/trades
type MarketTrades struct {
	Error  int      `json:"error"`
	Result [][4]any `json:"result"`
}

// /api/market/bids
type MarketBids struct {
	Error  int      `json:"error"`
	Result [][5]any `json:"result"`
}

// /api/market/asks
type MarketAsks struct {
	Error  int      `json:"error"`
	Result [][5]any `json:"result"`
}

// /api/market/books
type MarketBooks struct {
	Error  int `json:"error"`
	Result struct {
		Bids [][5]any `json:"bids"`
		Asks [][5]any `json:"asks"`
	} `json:"result"`
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
type MarketTicker struct {
	string struct {
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
}

// /api/market/symbols
type MarketSymbols struct {
	Error  int `json:"error"`
	Result []struct {
		ID     int    `json:"id"`
		Symbol string `json:"symbol"`
		Info   string `json:"info"`
	} `json:"result"`
}

// /api/v3/market/my-open-orders
type MyOpenOrder struct {
	Error  int `json:"error"`
	Result []struct {
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
	} `json:"result"`
}

// /api/v3/market/my-order-history
type MyOrderHistory struct {
	Error  int `json:"error"`
	Result []struct {
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
	} `json:"result"`
	Pagination struct {
		Page int `json:"page"`
		Last int `json:"last"`
		Next int `json:"next"`
		Prev int `json:"prev"`
	} `json:"pagination"`
}

// /api/v3/market/order-info
type OrderInfo struct {
	Error  int `json:"error"`
	Result struct {
		ID            string  `json:"id"`
		First         string  `json:"first"`
		Parent        string  `json:"parent"`
		Last          string  `json:"last"`
		Amount        string  `json:"amount"`
		Rate          int     `json:"rate"`
		Fee           int     `json:"fee"`
		Credit        int     `json:"credit"`
		Filled        float64 `json:"filled"`
		Total         int     `json:"total"`
		Status        string  `json:"status"`
		PartialFilled bool    `json:"partial_filled"`
		Remaining     int     `json:"remaining"`
		History       []struct {
			Amount    float64 `json:"amount"`
			Credit    float64 `json:"credit"`
			Fee       float64 `json:"fee"`
			Hash      string  `json:"hash"`
			ID        string  `json:"id"`
			Rate      int     `json:"rate"`
			Timestamp int64   `json:"timestamp"`
			TxnID     string  `json:"txn_id"`
		} `json:"history"`
	} `json:"result"`
}

type TradingCredit struct {
	Error  int     `json:"error"`
	Result float32 `json:"result"`
}

type Limits struct {
	Error  int `json:"error"`
	Result struct {
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
	} `json:"result"`
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

type Balance struct {
	Error  int                      `json:"error"`
	Result map[string]BalanceResult `json:"result"`
}

type BalanceResult struct {
	Available float64 `json:"available"`
	Reserved  float64 `json:"reserved"`
}

type WsToken struct {
	Error  int    `json:"error"`
	Result string `json:"result"`
}

type InternalWithdraw struct {
	Error  int `json:"error"`
	Result struct {
		Txn string  `json:"txn"`
		Adr string  `json:"adr"`
		Mem string  `json:"mem"`
		Cur string  `json:"cur"`
		Amt float64 `json:"amt"`
		Fee float64 `json:"fee"`
		Ts  int     `json:"ts"`
	} `json:"result"`
}

type DepositHistory struct {
	Error  int `json:"error"`
	Result []struct {
		Hash          string  `json:"hash"`
		Currency      string  `json:"currency"`
		Amount        float64 `json:"amount"`
		FromAddress   string  `json:"from_address"`
		ToAddress     string  `json:"to_address"`
		Confirmations int     `json:"confirmations"`
		Status        string  `json:"status"`
		Time          int     `json:"time"`
	} `json:"result"`
	Pagination struct {
		Page int `json:"page"`
		Last int `json:"last"`
	} `json:"pagination"`
}

type WithdrawHistory struct {
	Error  int `json:"error"`
	Result []struct {
		TxnID    string  `json:"txn_id"`
		Hash     string  `json:"hash"`
		Currency string  `json:"currency"`
		Amount   string  `json:"amount"`
		Fee      float64 `json:"fee"`
		Address  string  `json:"address"`
		Status   string  `json:"status"`
		Time     int     `json:"time"`
	} `json:"result"`
	Pagination struct {
		Page int `json:"page"`
		Last int `json:"last"`
	} `json:"pagination"`
}

type PlaceBid struct {
	Error  int `json:"error"`
	Result struct {
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
	} `json:"result"`
}

type PlaceAsk struct {
	Error  int `json:"error"`
	Result struct {
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
	} `json:"result"`
}

type CancelOrder struct {
	Error int `json:"error"`
}

type CryptoAddresses struct {
	Error  int `json:"error"`
	Result []struct {
		Currency string `json:"currency"`
		Address  string `json:"address"`
		Tag      int    `json:"tag"`
		Time     int    `json:"time"`
	} `json:"result"`
	Pagination struct {
		Page int `json:"page"`
		Last int `json:"last"`
	} `json:"pagination"`
}

type CryptoGenerateAddress struct {
	Error  int `json:"error"`
	Result []struct {
		Currency string `json:"currency"`
		Address  string `json:"address"`
		Memo     string `json:"memo"`
	} `json:"result"`
}

type CryptoWithdraw struct {
	Error  int `json:"error"`
	Result struct {
		Txn string  `json:"txn"`
		Adr string  `json:"adr"`
		Mem string  `json:"mem"`
		Cur string  `json:"cur"`
		Amt float64 `json:"amt"`
		Fee float64 `json:"fee"`
		Ts  int     `json:"ts"`
	} `json:"result"`
}
