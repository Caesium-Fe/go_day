package day7

import "fmt"

func Jiujiu() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v * %v = %2v  ", i, j, i*j)
		}
		fmt.Println()
	}
}
