package main

import "fmt"

func hammingDistance(x int, y int) int {
	var bits string
	var count int

	bits = fmt.Sprintf("%031b", x^y)
	for _, elem := range bits {
		if elem == 49 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(hammingDistance(93, 73))
}
