package bank

import (
	"bank/pkg/bank/types"
)

func PaymentSources(cards []types.Card) []types.PaymentSource {
	result:= make([]types.PaymentSource,3)
	var resultElement types.PaymentSource 
	for _,el:= range cards {
		if el.Active && el.Balance>0 {
			resultElement.Balance=el.Balance
			resultElement.Number=el.Number
			result = append(result, resultElement)
		}
	}
	return result
}
   