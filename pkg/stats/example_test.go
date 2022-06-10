package stats

import (
	"fmt"
	"github.com/bahodurnazarov/bankV2/v2/pkg/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{
			ID: 4047,
			Amount: 200,
		},
		{
			ID: 4041,
			Amount: 300,
		},
		
		{
			ID: 4045,
			Amount: 100,
		},
		}
	avg := Avg(payments)
	fmt.Println(avg)

	// Output: 200
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{
			ID: 4047,
			Amount: 200,
			Category: types.Category(types.StatusOk),
		},
		{
			ID: 4041,
			Amount: 300,
			Category: types.Category(types.StatusOk),
		},
		{
			ID: 4041,
			Amount: 300,
			Category: types.Category(types.StatusOk),
		},
		{
			ID: 4045,
			Amount: 100,
			Category: types.Category(types.StatusFail),
		},
		}
	total := TotalInCategory(payments, types.Category(types.StatusOk))
	fmt.Println(total)

	// Output: 800
}

