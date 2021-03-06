package stats

import (
	"fmt"
	"github.com/Iskandar2004/bank/pkg/types"
)


func ExampleAvg(){
	payments := []types.Payment{
		{
			ID:2,
			Amount:53_00,
			Category: "Cat",
		},
		{
			ID:1,
			Amount:51_00,
			Category: "Cat",
		},
		{
			ID:3,
			Amount:52_00,
			Category: "Cat",
		},
	}

	fmt.Println(Avg(payments))

	//Output: 5200
}

func ExampleTotalInCategory(){
	payments := []types.Payment{
		{
			ID:2,
			Amount:53_00,
			Category: "Cafe",
		},
		{
			ID:1,
			Amount:51_00,
			Category: "Cafe",
		},
		{
			ID:3,
			Amount:52_00,
			Category: "Restaurant",
		},
	}

	fmt.Println(TotalInCategory(payments, "Cafe"))

	//Output: 10400
}