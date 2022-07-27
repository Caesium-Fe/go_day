package day9

import (
	"fmt"
	"strings"
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
