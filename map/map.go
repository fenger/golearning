package _map

import (
	"fmt"
	"sort"
)

/*
Go语言中 map 是一种特殊的数据结构，一种元素对（pair）的无序集合，pair 对应一个 key（索引）和一个 value（值），
所以这个结构也称为关联数组或字典，这是一种能够快速寻找值的理想结构，给定 key，就可以迅速找到对应的 value。
map 这种数据结构在其他编程语言中也称为字典（Python）、hash 和 HashTable 等。

map 是引用类型，可以使用如下方式声明
```
var mapName map[keyType]valueType
```
其中：
- mapName 为 map 的变量名。
- keyType 为键类型。
- valueType 是键对应的值类型。

*/

func InitMap() {
	// 在声明的时候不需要知道 map 的长度，因为 map 是可以动态增长的，未初始化的 map 的值是 nil，使用函数 len() 可以获取 map 中 pair 的数目。
	var mapLit map[string]int
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}

	// 使用make创建map
	mapCreated := make(map[string]float32)

	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapCreated["key3"] = 3

	mapNames := make(map[string]string)
	mapNames["first"] = "美队"
	mapNames["second"] = "雷神"

	fmt.Printf("Map literal at \"one\" is : %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is : %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is : %d\n", mapAssigned["two"])
	fmt.Printf("Map literal at \"ten\" is : %d\n", mapLit["ten"])
	fmt.Printf("Map name at \"first\" is : %s\n", mapNames["first"])
	fmt.Printf("Map name at \"second\" is : %s\n", mapNames["third"])
}

// make(map[keyType]valueType, cap)
// 当 map 增长到容量上限的时候，如果再增加新的 key-value，map 的大小会自动加 1，所以出于性能的考虑，
// 对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。
func InitMapCap() {
	mapNames := make(map[string]string, 2)

	mapNames["first"] = "美队"
	mapNames["second"] = "雷神"
	mapNames["third"] = "黑寡妇"

	fmt.Printf("first value is : %s\n", mapNames["first"])
	fmt.Printf("third value is : %s\n", mapNames["third"])
}

// 将切片作为value值
func InitMapValueSlice() {
	mp := make(map[string][]string)

	names := []string{"美队", "雷神", "丹妮莉丝", "囧恩"}
	mp["game_of_thrones"] = names[2:]
	mp["avengers"] = names[:2]

	fmt.Printf("Game of Thrones contains %v\n", mp["game_of_thrones"])
	fmt.Printf("Avengers contains %v\n", mp["avengers"])
}

// 遍历map
// 遍历输出元素的顺序与定义时的顺序无关
func RangeMap() {
	week := make(map[string]int)
	week["SUNDAY"] = 0
	week["MONDAY"] = 1
	week["TUESDAY"] = 2

	for k, v := range week {
		// 无序
		fmt.Printf("key: %s, value: %d\n", k, v)
	}
}

// 特定顺序输出
func SortMap() {
	week := make(map[string]int)
	week["SUNDAY"] = 0
	week["MONDAY"] = 1
	week["TUESDAY"] = 2

	var weekSlice []string
	for k := range week {
		weekSlice = append(weekSlice, k)
	}

	// 切片排序
	sort.Strings(weekSlice)

	// 输出排序结果
	fmt.Println(weekSlice)
}

func DeleteItemFromMap() {
	jobs := make(map[string]string)
	jobs["golang"] = "nice"
	jobs["java"] = "good"
	jobs["rust"] = "better"

	fmt.Println(jobs)

	delete(jobs, "java")
	fmt.Println(jobs)
}
