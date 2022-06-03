package stats

import (
	"github.com/bahodurnazarov/bankV2/v2/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	for _, payment := range payments {
		if payment.Category != types.Category(types.StatusFail){
			sum += payment.Amount
		}
	}
	return sum / types.Money(len(payments))
}

func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, payment := range payments {
		if payment.Category != types.Category(types.StatusFail) {
			sum += payment.Amount
		}
	}
	return sum
}


// func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
// 	var sum types.Money
// 	for _, payment := range payments {
// 		if category != types.Category(types.StatusFail) {
// 			sum += payment.Amount
// 		}
// 	}
// 	return sum
// }