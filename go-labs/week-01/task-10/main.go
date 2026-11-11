package main 

import "fmt"

func main() {
	n := 4
	if n > 0 {
		fmt.Println("положительное")
	} else if n < 0 {
		fmt.Println("отрицательное")
	} else {
		fmt.Println("ноль")
	}
	
}