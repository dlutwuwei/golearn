package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}

	j := 0
I:
	j++
	fmt.Println(j)
	if j <= 10 {
		goto I
	}

	var list [10]int
	for k := range list {
		fmt.Println(k)
	}
}
