package main

import "fmt"

func imt(height float64, weight float64) float64 {
	return weight / (height * height)
}

func imtRes(res float64) {
	switch {
	case res < 18.5:
		fmt.Println("Дефицит массы тела")
	case res <= 24.9:
		fmt.Println("Нормальная масса тела")
	case res <= 29.9:
		fmt.Println("Увеличение массы тела")
	case res <= 34.9:
		fmt.Println("Ожирение 1 степени")
	case res <= 39.9:
		fmt.Println("Ожирение 2 степени")
	default:
		fmt.Println("Ожирение 3 степени")
	}
}

func main() {
	res := imt(1.78, 53.3)
	fmt.Println(res)
	imtRes(res)
}