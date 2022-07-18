package day7

import "fmt"

// type Person struct {
// 	age int
// }

func Test2() {
	person := &Person{28}

	// 1.28
	defer fmt.Println(person.age)

	// 2.28
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	// 3.29
	defer func() {
		fmt.Println(person.age)
	}()

	person = &Person{29}
}
