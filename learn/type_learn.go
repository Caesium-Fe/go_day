package learn

import (
	"fmt"
	"unsafe"
)

type Foo struct {
	A int8 // 1
	B int8 // 1
	C int8 // 1
}

type Bar3 struct {
	z bool  // 1
	x int32 // 4
	y *Foo  // 8
}

type student struct {
	name string
	age  int
}

func No_name_type() {

	// 结构体内存对齐问题
	var b3 Bar3
	fmt.Println(unsafe.Sizeof(b3)) // 16

	// 秘名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)
	fmt.Println(unsafe.Alignof(user))
	fmt.Println(unsafe.Alignof(user.Name))
	fmt.Println(unsafe.Alignof(user.Age))
	fmt.Println(unsafe.Sizeof(user))
	// for key, value := range user {
	// 	fmt.Println(key, value)
	// }

	// 面试题
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
		fmt.Println(&stu)
	}

	for k, v := range m {
		fmt.Println(v)
		fmt.Println(k, "=>", v.name)
	}
	// "小王子  =>  大王八"
	// "娜扎  =>  大王八"
	// "大王八  =>  大王八"
}

// type Person struct {
// 	name string
// 	city string
// 	age  int8
// }

// 构造函数
func newPerson(name, city string, age int8) *Person {
	return &Person{
		name: name,
		city: city,
		age:  age,
	}
}

//Dream Person做梦的方法 注意不是函数是方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func Use_init_func() {
	p9 := newPerson("张三", "武汉", 90)
	fmt.Printf("%#v\n", p9)
	//&main.person{name:"张三", city:"武汉", age:90}
	p1 := newPerson("小王子", "武汉", 25)
	p1.Dream()
}

//MyInt 将int定义为自定义MyInt类型
type MyInt int

//SayHello 为MyInt添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个int。")
}

func Try_type_function() {
	var m1 MyInt
	m1.SayHello() //Hello, 我是一个int。
	m1 = 100
	fmt.Printf("%#v  %T\n", m1, m1) //100  main.MyInt
}

// 继承的演示
//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func Demo_struct_inherit() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}
