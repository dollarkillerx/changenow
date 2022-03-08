package changenow

import "time"

type Currencie struct {
	Ticker            string `json:"ticker"`
	Name              string `json:"name"`
	Image             string `json:"image"`
	HasExternalId     bool   `json:"hasExternalId"`
	IsFiat            bool   `json:"isFiat"`
	Featured          bool   `json:"featured"`
	IsStable          bool   `json:"isStable"`
	SupportsFixedRate bool   `json:"supportsFixedRate"`
}

type MinAmount struct {
	MinAmount float64 `json:"minAmount"`
}

type TransactionsRequest struct {
	From          string `json:"from"`
	To            string `json:"to"`
	Address       string `json:"address"`
	Amount        string `json:"amount"`
	ExtraId       string `json:"extraId"`
	UserId        string `json:"userId"`
	ContactEmail  string `json:"contactEmail"`
	RefundAddress string `json:"refundAddress"`
	RefundExtraId string `json:"refundExtraId"`
}

type TransactionsResponse struct {
	PayinAddress  string  `json:"payinAddress"`
	PayoutAddress string  `json:"payoutAddress"`
	PayoutExtraId string  `json:"payoutExtraId"`
	FromCurrency  string  `json:"fromCurrency"`
	ToCurrency    string  `json:"toCurrency"`
	RefundAddress string  `json:"refundAddress"`
	RefundExtraId string  `json:"refundExtraId"`
	Id            string  `json:"id"`
	Amount        float64 `json:"amount"`
}

type TransactionStatus struct {
	Status                string    `json:"status"`
	PayinAddress          string    `json:"payinAddress"`
	PayoutAddress         string    `json:"payoutAddress"`
	FromCurrency          string    `json:"fromCurrency"`
	ToCurrency            string    `json:"toCurrency"`
	Id                    string    `json:"id"`
	UpdatedAt             time.Time `json:"updatedAt"`
	ExpectedSendAmount    int       `json:"expectedSendAmount"`
	ExpectedReceiveAmount float64   `json:"expectedReceiveAmount"`
	CreatedAt             time.Time `json:"createdAt"`
	IsPartner             bool      `json:"isPartner"`
}
