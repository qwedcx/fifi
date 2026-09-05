package main

import "fmt"

func numRes(num int) {
	var sum int
	sign := 1
	for i := 0; i < 6; i++ {
		if i == 3 {
			sign = -1
		}
		sum += num % 10 * sign
		num /= 10
	}
	if sum == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	var num int
	fmt.Scan(&num)
	numRes(num)
}
