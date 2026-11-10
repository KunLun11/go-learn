package main

import "fmt"

var number int = 42
var number2 int64 = 9000000000
var number3 float64 = 3.14
var name string = "go"
var boolname bool = true
var number4 uint = 65
var number5 int32 = 1071

func main() {
	fmt.Printf("number(%T)\n", number)
	fmt.Printf("number2(%T)\n", number2)
	fmt.Printf("number3(%T)\n", number3)
	fmt.Printf("name(%T)\n", name)
	fmt.Printf("boolname(%T)\n", boolname)
	fmt.Printf("number4(%T)\n", number4)
	fmt.Printf("number5(%T)\n", number5)
}