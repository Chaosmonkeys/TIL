package main

import "fmt"

var z int = 16

func main()  {
	var a int = 1
	fmt.Println(a)
	var b string = "hi"
	fmt.Println(b)
	c := 2
	fmt.Print(c)
	d := "bye"
	fmt.Println(d)

	fmt.Println(z)
}

func main()  {
	var a, b int = 1, 2
	fmt.Println(a, b)
	c, d, e := 3, 4, 5
	fmt.Println(c, d, e)
	var f, g string = "hi", "there"
	fmt.Println(f, g)
}