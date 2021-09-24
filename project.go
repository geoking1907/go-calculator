package main

import "fmt"

func main() {
	fmt.Print("Enter first number: ")
	var num1 float64
	fmt.Scanln(&num1)

	fmt.Print("Enter second number: ")
	var num2 float64
	fmt.Scanln(&num2)

	fmt.Print("Enter function: ")
	var f string
	fmt.Scanln(&f)
	var res float64
	switch f {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "*":
		res = num1 * num2
	case "/":
		res = num1 / num2
	default:
		fmt.Println("Error!")
		break
	}
	fmt.Print(res)
}
