package main

import (
	"fmt"
	"math"
)

type stuff struct {
	value  int
	weight int
	name   string
}

const len = 5

func main() {
	cake := []stuff{
		{
			weight: 2,
			value:  6,
			name:   "a",
		},
		{
			weight: 2,
			value:  3,
			name:   "b",
		},
		{
			weight: 6,
			value:  5,
			name:   "c",
		},
		{
			weight: 5,
			value:  4,
			name:   "d",
		},
		{
			weight: 4,
			value:  6,
			name:   "e",
		},
	}
	maxWeight := solve(5, 10, 0, cake)
	fmt.Println(maxWeight)

}

func solve(remainNum int, remainWeight int, curValue int, cake []stuff) (res int) {
	if remainNum == 1 {
		if cake[len-1].weight <= remainWeight {
			return cake[len-1].value + curValue
		} else {
			return curValue
		}
	}
	if remainWeight < cake[len-remainNum].weight {
		return solve(remainNum-1, remainWeight, curValue, cake)
	} else {
		res = int(math.Max(float64(solve(remainNum-1, remainWeight, curValue, cake)),
			float64(float64(solve(remainNum-1, remainWeight-cake[len-remainNum].weight, curValue+cake[len-remainNum].value, cake)))))
		return
	}
}
