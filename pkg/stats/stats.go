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

func FilterByCategory(payments []types.Payment, category types.Category) []types.Payment {
	var filtered []types.Payment
	for _, payment := range payments {
		if payment.Category == category {
			filtered = append(filtered, payment)
		}
	}
	return filtered
}


func CategoriesAvg(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}
	counter := map[types.Category]int64{}
	
	for _, payment := range payments {
		counter[payment.Category]++
		categories[payment.Category] += payment.Amount 
	}
	for k, v := range categories {
		categories[k] = types.Money(int64(v) / counter[k])
	}

	return categories 
}

func CategoriesTotal(payments []types.Payment) map[types.Category]types.Money {
	categories := map[types.Category]types.Money{}

	for _, payment := range payments {
		categories[payment.Category] += payment.Amount
	}

	return categories
}



