package learn

import (
	"fmt"
)

// 使用初始化列表来设置数组元素的值
func Shuzu_init_1() {
	//数组会初始化为int类型的零值
	var testArray [3]int
	//使用指定的初始值完成初始化
	var numArray = [3]int{1, 2}
	//使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"}
	fmt.Println(testArray) //[0 0 0]
	fmt.Println(numArray)  //[1 2 0]
	fmt.Println(cityArray) //[北京 上海 深圳]
}

// 让编译器根据初始值的个数自行推断数组的长度
func Shuzu_init_2() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳"}
	//[0 0 0]
	fmt.Println(testArray)
	//[1 2]
	fmt.Println(numArray)
	//type of numArray:[2]int
	fmt.Printf("type of numArray:%T\n", numArray)
	//[北京 上海 深圳]
	fmt.Println(cityArray)
	//type of cityArray:[3]string
	fmt.Printf("type of cityArray:%T\n", cityArray)
}

// 可以使用指定索引值的方式来初始化数组
func Shuzu_init_3() {
	a := [...]int{1: 1, 3: 5}
	// [0 1 0 5]
	fmt.Println(a)
	//type of a:[4]int
	fmt.Printf("type of a:%T\n", a)
}

// 遍历数组a有以下两种方法：
func Bianli() {
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}
}

// 二维数组
func Duowei() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	//[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a)
	//支持索引取值:重庆
	fmt.Println(a[2][1])
}

// 数组是值类型，赋值和传参会复制整个数组。
// 因此改变副本的值，不会改变本身的值。
// 注意：
// 数组支持 “==“、”!=” 操作符，因为内存总是被初始化过的。
// [n]*T表示指针数组，*[n]T表示数组指针 。
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}
func Modify() {
	a := [3]int{10, 20, 30}
	//在modify中修改的是a的副本x
	modifyArray(a)
	//[10 20 30]
	fmt.Println(a)
	b := [3][2]int{
		{1, 1},
		{1, 1},
		{1, 1},
	}
	//在modify中修改的是b的副本x
	modifyArray2(b)
	//[[1 1] [1 1] [1 1]]
	fmt.Println(b)
}

// Practice
// 求数组[1, 3, 5, 7, 8]所有元素的和
func Sumr() {
	arrayA := [...]int{1, 3, 5, 7, 8}
	lena := len(arrayA)
	sumr := 0
	for i := 0; i < lena; i++ {
		sumr += arrayA[i]
	}
	fmt.Printf("reasult is %d\n", sumr)
}

// 找出数组中和为指定值的两个元素的下标，比如从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。
func Find_two_index() {
	// aim := bufio.NewReader(os.Stdin)
	arrayA := [...]int{1, 3, 5, 7, 8}
	lena := len(arrayA)
	var aim int
	fmt.Scan(&aim)
	for i := 0; i < lena; i++ {
		for j := i; j < lena; j++ {
			if aim == arrayA[i]+arrayA[j] {
				fmt.Printf("%d , %d", arrayA[i], arrayA[j])
			}
		}
	}
}
