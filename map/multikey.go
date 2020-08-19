package _map

import "fmt"

// 查询键
type queryKey struct {
	Name string
	Age  int
}

func RunMultiKeyIndex2() {
	list := []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 21},
		{Name: "王五", Age: 22},
	}

	buildIndex2(list)

	queryData2("张三", 30)
}

var mp2 = make(map[queryKey]*Profile)

func buildIndex2(list []*Profile) {
	for _, profile := range list {
		key := queryKey{
			Name: profile.Name,
			Age:  profile.Age,
		}

		mp2[key] = profile
	}
}

func queryData2(name string, age int) {
	key := queryKey{name, age}

	result, ok := mp2[key]
	if ok {
		fmt.Println(result)
	} else {
		fmt.Println("not found")
	}
}
