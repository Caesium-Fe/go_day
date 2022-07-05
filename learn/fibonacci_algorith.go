package learn

import "fmt"

var fibonaccires []int

func fibonacci(n int) int {
	if n<3{
		return 1
	}
	if fibonaccires[n] == 0{
		fibonaccires[n] = fibonacci(n-2) + fibonacci(n-1)
	}
	return fibonaccires[n]
}

func Recursion() {
	n := 45
	fibonaccires = make([]int, n+1)
	c := fibonacci(n)
	fmt.Println(c)
}
