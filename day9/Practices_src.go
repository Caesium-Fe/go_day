package day9

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func get_count(name string) int {
	var total, count int
	get_coin_arr := []string{"e", "E", "i", "I", "o", "O", "u", "U"}
	for _, char := range get_coin_arr {
		count = strings.Count(name, char)
		switch char {
		case "i", "I":
			count *= 2
		case "o", "O":
			count *= 3
		case "u", "U":
			count *= 4
		default:
			fmt.Println("the name is great")
		}
		total += count
	}
	return total
}

func dispatchCoin() int {
	var everyTotal, eachTotal int
	for _, name := range users {
		// 获得当前名字人获得金币数量
		eachTotal = get_count(name)
		// 生成[name:num_coin]
		distribution[name] = eachTotal
		// 金币总数
		everyTotal += eachTotal

		// go func(x string) {
		// 	v := strings.Count(name, x)
		// 	c <- v
		// }("e")
		// go func(x string) {
		// 	v := strings.Count(name, x)
		// 	c <- v
		// }("E")
		// go func(x string) {
		// 	v := strings.Count(name, x)
		// 	c <- v
		// }("i")
		// go func(x string) {
		// 	v := strings.Count(name, x)
		// 	c <- v
		// }("I")
		// go func(x string) {
		// 	v := strings.Count(name, x)
		// 	c <- v
		// }("o")
		// go func(x string) {
		// 	v := strings.Count(name, x)
		// 	c <- v
		// }("O")
		// go func(x string) {
		// 	v := strings.Count(name, x)
		// 	c <- v
		// }("u")
		// go func(x string) {
		// 	v := strings.Count(name, x)
		// 	c <- v
		// }("U")
		// close(c)
		// <-done
		// go func() {
		// 	defer close(done)
		// 	eachTotal = 0
		// 	for {
		// 		x, ok := <-c
		// 		// if !ok { //判断通道是否关闭
		// 		// 	return
		// 		// }
		// 		// println(x)
		// 		// eachTotal = eachTotal + x
		// 		eachTotal += x
		// 	}
		// }()

		// strings.Count(name, "")
	}
	// 余下金币数量
	return coins - everyTotal
}

func Dispatch_golden_coin() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

/*
1使用“面向对象”的思维方式编写一个学生信息管理系统。
	1学生有id、姓名、年龄、分数等信息
	2程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
*/
type Student1 struct {
	ID    int
	Name  string
	Age   int
	Score int
}

func creat_students() []*Student1 {
	Students := make([]*Student1, 0, 200)
	for i := 0; i < 10; i++ {
		stu := &Student1{
			Name:  fmt.Sprintf("stu%02d", i),
			Age:   i,
			ID:    i,
			Score: i,
		}
		Students = append(Students, stu)
	}

	return Students
}

// 方法展示学生信息
func (s *Student1) show_students() {
	fmt.Println(s.ID, s.Name, s.Age, s.Score)
}

func Struct_of_studens() {
	// 所有学生
	Students := creat_students()
	// 获取学生信息
	for _, k := range Students {
		k.show_students()
	}

	// 添加学生

	// 修改学生信息

	//删除学生

}

//使用 goroutine 和 channel 实现一个计算int64随机数各位数和的程序，例如生成随机数61345，计算其每个位数上的数字之和为19。

var wg11 sync.WaitGroup

func Test_of_sum() {
	wg11.Add(1)
	jobChan := make(chan int64, 1)
	resultChan := make(chan int64, 100)
	go randNum(jobChan)
	for i := 0; i < 24; i++ {
		wg11.Add(1)
		go sumNum(jobChan, resultChan)
	}
	// 发现wg.Wait()要放在resultChan关闭之前
	//不然会报“往关闭的通道里写入数据的错误"
	wg11.Wait()
	close(resultChan)
	// i := 0
	for ret := range resultChan {
		fmt.Println(ret)
		// i++
	}
	// fmt.Println(i)
}

func randNum(jobChan chan<- int64) {
	defer wg11.Done()
	for i := 0; i < 100; i++ {
		jobChan <- rand.Int63()
	}
	close(jobChan)
}

func sumNum(jobChan <-chan int64, resultChan chan<- int64) {
	defer wg11.Done()
	for v := range jobChan {
		var s int64
		for v > 0 {
			s += v % 10
			v = v / 10
		}
		resultChan <- s
	}
}

// // 开启一个 goroutine 循环生成int64类型的随机数，发送到jobChan
// func creat_randnum(){
// 	var ranumb := make(chan int64)
// 	defer wg11.Done()
// 	wg11.Add(1)
// 	go func(){
// 		for i:=0;i<100;i++{
// 			r:=rand.New(rand.NewSource(time.Now().UnixNano()))
// 			ranumb<-r.Int63()
// 		}
// 		close(ranumb)
// 	}()
// 	// return ranumb
// }

// // 开启24个 goroutine 从jobChan中取出随机数计算各位数的和，将结果发送到resultChan
// func cal_sum(){

// }

// // 主 goroutine 从resultChan取出结果并打印到终端输出
