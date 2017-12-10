package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{3, 4, 5}, 9))

}

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for k, v := range nums {
		another := dict[target-v]
		if dict[v] == 0 {
			dict[v] = k + 1
		}
		if another != 0 {
			return []int{another - 1, k}
		}
	}
	return []int{}
}
