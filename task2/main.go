package main

import "fmt"

func numRes(num int) {
	if num%10 != (num/10)%10 && num%10 != (num/100)%10 && (num/10)%10 != (num/100)%10 {
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
