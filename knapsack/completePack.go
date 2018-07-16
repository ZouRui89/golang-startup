package main

import (
	"fmt"
	"math"
)

type stuff2 struct {
	value  int
	weight int
	name   string
}

const len2 = 5

func main() {
	cake := []stuff2{
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
	maxWeight := solveCom(5, 10, 0, cake)
	fmt.Println(maxWeight)

}

func solveCom(remainNum int, remainWeight int, curValue int, cake []stuff2) (res int) {
	if remainNum == 1 {
		if cake[len2-1].weight <= remainWeight {
			return cake[len2-1].value + curValue
		} else {
			return curValue
		}
	}
	if remainWeight < cake[len2-remainNum].weight {
		return solveCom(remainNum-1, remainWeight, curValue, cake)
	} else {
		res = int(math.Max(float64(solveCom(remainNum-1, remainWeight, curValue, cake)),
			float64(float64(solveCom(remainNum, remainWeight-cake[len2-remainNum].weight, curValue+cake[len2-remainNum].value, cake)))))
		return
	}
}
