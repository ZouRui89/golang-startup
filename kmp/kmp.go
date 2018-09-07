package main

import (
	"fmt"
)

func cal_next(ptr string) [10]int {
	pLen := len(ptr)
	k := -1
	var next [10]int

	for i := 1; i < pLen; i++ {
		for k > -1 && ptr[k+1] != ptr[i] {
			k = next[k]
		}
		if ptr[k+1] == ptr[i] {
			k = k + 1
		}
		next[i] = k
	}
	return next

}
func main() {
	str := "abcdabccabcdabababcd"
	ptr := "abcd"

	next := cal_next(ptr)

	k := -1
	for i := 0; i < len(str); i++ {
		for k > -1 && ptr[k+1] != str[i] {
			k = next[k]
		}
		if ptr[k+1] == str[i] {
			k = k + 1
		}
		if k == len(ptr)-1 {
			fmt.Println(i - len(ptr) + 1)
			k = -1
			i = i - len(ptr) + 1
		}
	}
}
