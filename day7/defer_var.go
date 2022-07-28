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

// 经典测试
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() int {
	x := 5
	defer func(x int) {
		// fmt.Println("sad %v", x)
		x++
		fmt.Println("sad %v", &x)
	}(x)
	fmt.Println("sad %v", &x)
	return x
}

func f3() (y int) {
	x := 5
	defer func() {
		fmt.Println("sad %v", y)
		y++
		fmt.Println("sad %v", x)
		fmt.Println("sad %v", y)
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		fmt.Println("asd %v", x)
		x++
		fmt.Println("asd %v", x)
	}(x)
	fmt.Println("asd %v", x)
	return 5
}
func Test3() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

// defer面试题
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func Test4() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

func Defer_interface() {
	fmt.Println("main", inc())
	// result:
	// test 3 1
	// test 2 2
	// inc 2
	// test 1 3
	// main 2
}

func inc() int {
	t := &test{num: 0}
	defer t.Inc(3).Inc(2).Inc(1)
	fmt.Println("inc", t.num)
	return t.num
}

type test struct {
	num int
}

func (t *test) Inc(flag int) *test {
	t.num++
	fmt.Println("test", flag, t.num)
	return t
}
