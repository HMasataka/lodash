package main

import (
	"fmt"

	"github.com/samber/lo"
)

func main() {
	even := lo.Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})
	fmt.Println(even)

	present := lo.ContainsBy([]int{0, 1, 2, 3, 4, 5}, func(x int) bool {
		return x == 3
	})
	fmt.Println(present)

	lo.ForEach([]string{"hello", "world"}, func(x string, _ int) {
		fmt.Println(x)
	})
}
