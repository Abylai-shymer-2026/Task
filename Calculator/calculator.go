package main

import "fmt"

type Cal struct{}

func (k Cal) Add(a, b float64) float64 {
	return a + b
}

func (k Cal) Subtract(a, b float64) float64 {
	return a - b
}

func (k Cal) Multi(a, b float64) float64 {
	return a * b
}

func (k Cal) Divide(a, b float64) float64 {
	return a / b
}

func main() {
	calculator := Cal{}
	var a, b float64
	fmt.Scan(&a, &b)
	ans := calculator.Add(a, b)
	fmt.Println(ans)

	ans = calculator.Subtract(a, b)
	fmt.Println(ans)

	ans = calculator.Multi(a, b)
	fmt.Println(ans)

	ans = calculator.Divide(a, b)
	fmt.Println(ans)
}
