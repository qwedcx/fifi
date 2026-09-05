package main

import "fmt"

func main() {
	var num string
	fmt.Scan(&num)
	digit := int(num[0] - '0')
	fmt.Println(digit)
}
