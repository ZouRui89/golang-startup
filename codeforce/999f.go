/*
	http://codeforces.com/contest/999/problem/F
	F. Cards and Joy
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var playerNum, evenCardNum int
	fmt.Scanln(&playerNum, &evenCardNum)

	var tmp int
	var cardCount [100005]int
	for i := 0; i < playerNum*evenCardNum; i++ {
		fmt.Scanf("%d", &tmp)
		cardCount[tmp]++
	}

	fmt.Scanf("%d", &tmp)
	var favCount [100005]int
	for i := 0; i < playerNum; i++ {
		fmt.Scanf("%d", &tmp)
		favCount[tmp]++
	}

	fmt.Scanf("%d", &tmp)
	var joy [15]int
	for i := 0; i < evenCardNum; i++ {
		fmt.Scanf("%d", &tmp)
		joy[i] = tmp
	}

	// distribute i cards(same card number) among j players
	var dp [5005][505]int
	for i := 1; i <= playerNum*evenCardNum; i++ {
		dp[i][1] = joy[int(math.Min(float64(i), float64(evenCardNum)))-1]
		for j := 2; j <= playerNum; j++ {
			for k := 1; k <= int(math.Min(float64(i), float64(evenCardNum))); k++ {
				dp[i][j] = int(math.Max(float64(dp[i][j]), float64(dp[i-k][j-1]+joy[k-1])))
			}
		}
	}

	var res int64
	for i := 1; i <= 100001; i++ {
		if cardCount[i] > 0 {
			res = res + int64(dp[cardCount[i]][favCount[i]])
		}
	}
	fmt.Println(res)
}
