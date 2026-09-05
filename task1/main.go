package main

import "fmt"

func intRes(num int) {
	switch {
	case num > 0:
		fmt.Println("Число положительное")
	case num < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("Ноль")
	}
}

func main() {
	var num int
	fmt.Scan(&num)
	intRes(num)
}
