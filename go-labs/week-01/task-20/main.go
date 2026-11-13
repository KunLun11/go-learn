package main 

import "fmt"

func main() {
	a, b := 10.0, 4.0
	op := "/"
	switch {
	case op == "+":
		result := a + b
		fmt.Println(result)
	case op == "-":
		result := a - b
		fmt.Println(result)
	case op == "*":
		result := a * b
		fmt.Println(result)
	case op == "/":
		result := a / b
		fmt.Println(result)
	default:
		fmt.Println("Некорректный оператор")	
	}
}