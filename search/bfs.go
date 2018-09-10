package main

import (
	"container/list"
	"fmt"
)

var b [11][11]int
var hasVisted [11]bool
var cityN, roadN int

func main() {
	fmt.Scanf("%d", &cityN)
	fmt.Scanf("%d", &roadN)
	init_bfs_graph(roadN)

	bfs(1)

}

func init_bfs_graph(roadNum int) {
	var x, y int
	for i := 0; i < roadNum; i++ {
		fmt.Scanf("%d %d", &x, &y)
		b[x][y] = 1
		b[y][x] = 1
	}
}

func bfs(n int) {
	list := list.New()
	hasVisted[n] = true
	list.PushBack(n)
	fmt.Printf("%d ", n)

	for list.Len() > 0 {
		f := list.Front()
		list.Remove(f)

		for i := 0; i <= cityN; i++ {
			if b[f.Value.(int)][i] == 1 && hasVisted[i] == false {
				hasVisted[i] = true
				list.PushBack(i)
				fmt.Printf("%d ", i)
			}
		}
	}

}
