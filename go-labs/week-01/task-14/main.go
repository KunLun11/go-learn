package main 

import "fmt"

func main() {
	a, b, c := 5, 5, 8
	if a < b + c && b < a + c && c < a + b && a>0 && b>0 && c>0 {
		fmt.Println("Треугольник существует")
	} else {
		fmt.Println("Треугольник не существует")
	}
	switch {
		case a==b && b==c:
			fmt.Println("Треугольник равносторонний")
		case a==b || b==c || a==c:
			fmt.Println("Треугольник равнобедренный")
		default:
			fmt.Println("Треугольник разносторонний")
	}
}