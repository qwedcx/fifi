package main

import "fmt"

func intRes() string{
switch {
	case  > 0:
		fmt.Println("Число положительное")
	case  < 0:
		fmt.Println("Число отрицательное")
	default:
		fmt.Println("Число равно 0")
}
}


func main() {
    var name int
    fmt.Print("Введите целое число: ")
    Println(Scan(&name) )
}
