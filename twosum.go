package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for k, v := range nums {
		for i := k + 1; i < len(nums); i++ {
			if target-v == nums[i] {
				return []int{k, i}
			}
		}
	}
	return []int{}
}

func main() {
	a := []int{1, 5, 4}
	res := twoSum(a, 7)
	fmt.Println(res)
}
