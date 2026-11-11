package main 

import "fmt"

func main() {
	year := 2040  
	if year%4 == 0 && year%100 != 0 || year%400==0 {
		fmt.Println("Високосный год")
	} else {
		fmt.Println("Не високосный год")
	}
}
