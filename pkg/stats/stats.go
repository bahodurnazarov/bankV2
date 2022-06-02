package stats

import (
	"github.com/bahodurnazarov/goTask/pkg/types"
)

func Avg(payments []types.Payment) types.Money {
	var sum types.Money
	for _, payment := range payments {
		sum += payment.Amount
	}
	return sum / types.Money(len(payments))
}


func TotalInCategory(payments []types.Payment, category types.Category) types.Money {
	var sum types.Money
	for _, payment := range payments {
		if payment.Category == category {
			sum += payment.Amount
		}
	}
	return sum
}