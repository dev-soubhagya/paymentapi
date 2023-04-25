package services

type Request struct {
	Userid     int64   `json:"userid"`
	Merchantid int64   `json:"merchantid"`
	Amount     float64 `json:"amount"`
}
type Response struct {
	TxnId       int64  `json:"txnid"`
	Status      string `json:"status"`
	Status_desc string `json:"status_desc"`
}
type User struct {
	Userid  int64   `json:"userid"`
	Balance float64 `json:"balance"`
	Status  int     `json:"status"`
}
type Merchant struct {
	Merchantid int64   `json:"merchantid"`
	Balance    float64 `json:"balance"`
	Status     int     `json:"status"`
}
type User_transaction_status struct {
	Txnid   int64   `json:"txnid"`
	Userid  int64   `json:"userid"`
	Amount  float64 `json:"amount"`
	Balance float64 `json:"balance"`
	Status  string  `json:"status"`
	Type    int     `json:"type"`
}
type Merchant_transaction_status struct {
	Txnid      int64   `json:"txnid"`
	Merchantid int64   `json:"mechantid"`
	Amount     float64 `json:"amount"`
	Balance    float64 `json:"balance"`
	Status     string  `json:"status"`
	Type       string  `json:"type"`
}
