package day7

import "fmt"

type Person struct {
	age int
}

func Test1() {
	person := &Person{28}

	// 1.28
	defer fmt.Println(person.age)

	// 2.29
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.29
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}
