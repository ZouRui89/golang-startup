package main

import (
	"fmt"
)

var a [11][11]int
var visted [11]bool
var cityNum, roadNum int

func main() {
	fmt.Scanf("%d", &cityNum)
	fmt.Scanf("%d", &roadNum)
	init_graph(roadNum)
	fmt.Println(a)
	dfs(1)
}

func init_graph(roadNum int) {
	var x, y int
	for i := 0; i < roadNum; i++ {
		fmt.Scanf("%d %d", &x, &y)
		a[x][y] = 1
		a[y][x] = 1
	}
}

func dfs(u int) {
	fmt.Printf("%d ", u)
	visted[u] = true

	for i := 0; i <= cityNum; i++ {
		if a[u][i] == 1 && visted[i] == false {
			dfs(i)
		}
	}

}
