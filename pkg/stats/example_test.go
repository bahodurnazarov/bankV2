package stats

import (
	"fmt"
	"github.com/bahodurnazarov/goTask/pkg/types"
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
			Category: "Food",
		},
		{
			ID: 4043,
			Amount: 100,
			Category: "Car",
		},
		{
			ID: 4041,
			Amount: 300,
			Category: "Car",
		},
		
		{
			ID: 4045,
			Amount: 100,
			Category: "Car",
		},
		}
	tic := TotalInCategory(payments, "Car")
	fmt.Println(tic)

	// Output: 500
}
