package main

import "fmt"

func main() {
	fmt.Println(judgeCircle("LR"))
}

func judgeCircle(moves string) bool {
	var up, right, down, left int
	for _, elem := range moves {
		var move = string(elem)
		switch move {
		case "L":
			left++
		case "U":
			up++
		case "R":
			right++
		case "D":
			down++
		}
	}
	if right == left && up == down {
		return true
	} else {
		return false
	}
}
