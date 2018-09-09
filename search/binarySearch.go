package main

import (
	"fmt"
	"sort"
)

func main() {
	var slice []int
	slice = make([]int, 0)
	var n int
	var err error
	for {
		_, err = fmt.Scanf("%d", &n)
		if err != nil {
			break
		}
		slice = append(slice, n)
	}
	//fmt.Println(slice)
	sort.Sort(sort.IntSlice(slice))
	var value int
	fmt.Scanf("%d", &value)
	pos := binarySearch(slice, value)
	fmt.Printf("position %d", pos)
}

func binarySearch(slice []int, value int) int {
	len := len(slice)
	var low, high, mid int
	high = len - 1
	for low <= high {
		mid = (low + high) / 2
		if slice[mid] == value {
			return mid + 1
		} else {
			if slice[mid] > value {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return -1

}
