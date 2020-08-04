package main

import "fmt"

func main()  {
	const a int = 1 // 첫번째 값은 무조건 초기화
	const b, d int = 2, 3
	const c = "hi there" // 다른 형 선언 가능
	const (
		f = 55
		k = 44
		u
		r
		t = "end"
	)

	fmt.Println(f, k, u , r, t)
}
