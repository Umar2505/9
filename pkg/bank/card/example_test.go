package bank

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExamplePaymentSources_positive() {
	cards := []types.Card{
		{
			Balance: 1000,
			Number: "5055 xxxx xxxx 1233",
			Active: true,
		},
		{
			Balance: 5000,
			Number: "5055 xxxx xxxx 4233",
			Active: true,
		},
		{
			Balance: 100000,
			Number: "5035 xxxx xxxx 1233",
			Active: false,
		},
	}
	result := PaymentSources(cards)
	for _, v := range result {
		fmt.Println(v.Number)
	}
	//Output:
	//5055 xxxx xxxx 1233
	//5055 xxxx xxxx 4233
}