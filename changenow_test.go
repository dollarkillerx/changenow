package changenow

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestChangenowSDK_Currencies(t *testing.T) {
	sdk := New("")
	currencies, err := sdk.Currencies()
	if err != nil {
		panic(err)
	}

	print(currencies)
}

func TestChangenowSDK_MinAmount(t *testing.T) {
	sdk := New("")
	minAmount, err := sdk.MinAmount("usdttrc20_xmr")
	if err != nil {
		panic(err)
	}

	print(minAmount)
}

func TestChangenowSDK_Transactions(t *testing.T) {
	sdk := New("")
	transactions, err := sdk.Transactions(TransactionsRequest{
		From:    "usdttrc20",
		To:      "xmr",
		Address: "44AFFq5kSiGBoZ4NMDwYtN18obc8AemS33DBLWs3H7otXft3XjrpDtQGv7SqSsaBYBb98uNbr2VBBEt7f2wfn3RVGQBEP3A",
		Amount:  "10",
	})
	if err != nil {
		panic(err)
	}

	print(transactions)

	/**
	{
	    "payinAddress":"TSASMwv6b7n3xTgeBFsmUepmp3BvsG7boB",
	    "payoutAddress":"44AFFq5kSiGBoZ4NMDwYtN18obc8AemS33DBLWs3H7otXft3XjrpDtQGv7SqSsaBYBb98uNbr2VBBEt7f2wfn3RVGQBEP3A",
	    "payoutExtraId":"",
	    "fromCurrency":"usdttrc20",
	    "toCurrency":"xmr",
	    "refundAddress":"",
	    "refundExtraId":"",
	    "id":"28658d18aa648e",
	    "amount":0.0568745
	}
	*/
}

func TestChangenowSDK_TransactionStatus(t *testing.T) {
	sdk := New("")
	status, err := sdk.TransactionStatus("28658d18aa648e")
	if err != nil {
		panic(err)
	}

	print(status)

	/**
	{
	    "status":"waiting",
	    "payinAddress":"TSASMwv6b7n3xTgeBFsmUepmp3BvsG7boB",
	    "payoutAddress":"44AFFq5kSiGBoZ4NMDwYtN18obc8AemS33DBLWs3H7otXft3XjrpDtQGv7SqSsaBYBb98uNbr2VBBEt7f2wfn3RVGQBEP3A",
	    "fromCurrency":"usdttrc20",
	    "toCurrency":"xmr",
	    "id":"28658d18aa648e",
	    "updatedAt":"2022-03-08T03:17:57.381Z",
	    "expectedSendAmount":10,
	    "expectedReceiveAmount":0.0568745,
	    "createdAt":"2022-03-08T03:17:57.381Z",
	    "isPartner":false
	}

	Transaction status:
	new,
	waiting,
	confirming,
	exchanging,
	sending,
	finished,
	failed,
	refunded,
	verifying
	Transaction status:
	new,
	waiting,
	confirming,
	exchanging,
	sending,
	finished,
	failed,
	refunded,
	verifying
	*/
}

func print(r interface{}) {
	marshal, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(marshal))
}
