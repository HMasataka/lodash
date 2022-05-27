package main

import (
	"fmt"
	"strconv"

	"github.com/samber/lo"
)

func main() {
	even := lo.Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})
	fmt.Println(even)

	present := lo.Contains([]int{0, 1, 2, 3, 4, 5}, 5)
	fmt.Println(present)

	lo.ForEach([]string{"hello", "world"}, func(x string, _ int) {
		fmt.Println(x)
	})

	sum := lo.Reduce([]int{1, 2, 3, 4}, func(agg int, item int, _ int) int {
		return agg + item
	}, 0)
	fmt.Println(sum)

	times := lo.Times(3, func(i int) string {
		return strconv.FormatInt(int64(i), 10)
	})
	fmt.Println(times)

	uniq := lo.Uniq([]int{1, 2, 2, 1})
	fmt.Println(uniq)

	groups := lo.GroupBy([]int{0, 1, 2, 3, 4, 5}, func(i int) int {
		return i % 3
	})
	fmt.Println(groups)

	chunk := lo.Chunk([]int{0, 1, 2, 3, 4, 5}, 2)
	fmt.Println(chunk)

	partitions := lo.PartitionBy([]int{-2, -1, 0, 1, 2, 3, 4, 5}, func(x int) string {
		if x < 0 {
			return "negative"
		} else if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	fmt.Println(partitions)

	randomOrder := lo.Shuffle([]int{0, 1, 2, 3, 4, 5})
	fmt.Println(randomOrder)

}
