package learn

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// 统计句中单词出现的频率
func test1_cishu(example string) {
	var totalstr = make(map[string]int, 5)
	c := strings.Split(example, " ")
	fmt.Println(len(c))
	for index, value := range c {
		fmt.Println(index, value)
		totalstr[value] = strings.Count(example, value)
	}
	fmt.Println(totalstr)
	// n := strings.Count(example, "how")
	// fmt.Println(n)
}

// 下述输出结果为什么
func test2_print() {
	type Map map[string][]int
	m := make(Map)
	s := []int{1, 2}
	s = append(s, 3)
	fmt.Printf("%+v\n", s) // 1 2 3
	m["q1mi"] = s
	s = append(s[:1], s[2:]...)
	fmt.Printf("%+v\n", s) // 1 3
	fmt.Println(&s[0])
	fmt.Printf("%+v\n", m["q1mi"]) // 1 3 3
	fmt.Println(&m["q1mi"][0])
}

func map_slice() {
	var mapSlice = make([]map[string]string, 3)
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	// 对切片中的map元素进行初始化
	mapSlice[0] = make(map[string]string, 10)
	mapSlice[0]["name"] = "小王子"
	mapSlice[0]["password"] = "123456"
	mapSlice[0]["address"] = "沙河"
	for index, value := range mapSlice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
}

func map_slice_type() {
	var sliceMap = make(map[string][]string, 3)
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "中国"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "北京", "上海")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}

func map_slice_sort() {
	rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机整数
		scoreMap[key] = value
	}
	//取出map中的所有key存入切片keys
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}
	//对切片进行排序
	sort.Strings(keys)
	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}
}

func Map_all_test() {
	test2_print()
	var example string
	example = "how do you do"
	test1_cishu(example)
	map_slice()
	map_slice_type()
	map_slice_sort()
}
