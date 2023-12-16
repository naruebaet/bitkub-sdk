package response

// /api/status
type Status []struct {
	Name    string `json:"name"`
	Status  string `json:"status"`
	Message string `json:"message"`
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
