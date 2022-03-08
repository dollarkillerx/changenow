package changenow

import (
	"fmt"
	"time"

	"github.com/dollarkillerx/urllib"
)

type ChangenowSDK struct {
	apiKey  string
	apiUrl  string
	timeout time.Duration
}

// New https://changenow.io/affiliate
func New(apiKey string) *ChangenowSDK {
	if apiKey == "" {
		apiKey = "b0dd7347cf5448d6779bec284355bc61bb74daf18c513777bc98456b9ab36117"
	}
	return &ChangenowSDK{
		apiKey:  apiKey,
		apiUrl:  "https://api.changenow.io",
		timeout: time.Second * 3,
	}
}

func (c *ChangenowSDK) SetAPIUrl(url string) *ChangenowSDK {
	c.apiUrl = url
	return c
}

func (c *ChangenowSDK) SetTimeout(timeout time.Duration) *ChangenowSDK {
	c.timeout = timeout
	return c
}

// Currencies https://documenter.getpostman.com/view/8180765/SVfTPnM8?version=latest#f5216aba-6a44-49eb-a075-ad4435aa40db
func (c *ChangenowSDK) Currencies() ([]Currencie, error) {
	url := fmt.Sprintf("%s/v1/currencies?active=true&fixedRate=true", c.apiUrl)
	var currencies []Currencie
	err := urllib.Get(url).SetTimeout(c.timeout).KeepAlives().FromJson(&currencies)
	return currencies, err
}

// MinAmount https://documenter.getpostman.com/view/8180765/SVfTPnM8?version=latest#6f50a577-8558-4f58-b90c-85dfaab42c81
func (c *ChangenowSDK) MinAmount(fromTo string) (*MinAmount, error) {
	url := fmt.Sprintf("%s/v1/min-amount/%s?api_key=%s", c.apiUrl, fromTo, c.apiKey)
	var minAmount MinAmount
	err := urllib.Get(url).SetTimeout(c.timeout).KeepAlives().FromJson(&minAmount)
	return &minAmount, err
}

// Transactions https://documenter.getpostman.com/view/8180765/SVfTPnM8?version=latest#dfe05b67-8453-462e-b4dd-fa4b0001c197
func (c *ChangenowSDK) Transactions(req TransactionsRequest) (*TransactionsResponse, error) {
	url := fmt.Sprintf("%s/v1/transactions/%s", c.apiUrl, c.apiKey)
	var resp TransactionsResponse
	err := urllib.Post(url).SetJsonObject(req).SetTimeout(c.timeout).KeepAlives().FromJson(&resp)
	return &resp, err
}

// TransactionStatus https://documenter.getpostman.com/view/8180765/SVfTPnM8?version=latest#fa12244b-f879-4675-a6f7-553cc59435dc
func (c *ChangenowSDK) TransactionStatus(transactionsID string) (*TransactionStatus, error) {
	url := fmt.Sprintf("%s/v1/transactions/%s/%s", c.apiUrl, transactionsID, c.apiKey)
	var transactionStatus TransactionStatus
	err := urllib.Get(url).SetTimeout(c.timeout).KeepAlives().FromJson(&transactionStatus)
	return &transactionStatus, err
}
