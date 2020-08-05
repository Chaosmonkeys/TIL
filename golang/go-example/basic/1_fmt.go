package main

import "fmt"

func main() {
	var num1 int = 1
	var num2 int = 2

	fmt.Print("Hello goorm!", num1, num2, "\n")
	fmt.Println("Hi there", num1, num2)
	fmt.Printf("num1의 값:%d\nnum2의 값:%d\n", num1, num2)
	fmt.Printf("num1의 값:%d\nnum2의 값:%d\n", num1, num2)

	// /n : 개행, %d : 정수, %s : 문자 - Printf
}