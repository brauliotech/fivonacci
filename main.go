package main

import (
	"flag"
	"fmt"
)

var (
	list []int
)

func main() {
	limit := flag.Int("limit", 5, "Get position of serie")
	flag.Parse()

	fc := fibonacci()
	for i := 0; i < *limit; i++ {
		fmt.Println(fc(i))
	}
}

func fibonacci() func(int) int {
	return func(x int) int {
		if x >= 2 {
			xf := list[x-1] + list[x-2]
			list = append(list, xf)
			return (list[x])
		}

		list = append(list, x)
		return (list[x])
	}
}
