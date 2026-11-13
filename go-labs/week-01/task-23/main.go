package main

import "fmt"

func main() {
	password := "golang1234123gogogo"
	switch n := len(password);{
	case n < 8:
		fmt.Println("Слабый пароль")
	case n > 8 && n < 12:
		fmt.Println("Средний пароль")
	default:
		fmt.Println("Сильный")
	}
}
