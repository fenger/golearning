package _map

import "fmt"

type Profile struct {
	Name    string
	Age     int
	Married bool
}

// map多键索引
func RunMultiKeyIndex() {
	list := []*Profile{
		{Name: "张三", Age: 30, Married: true},
		{Name: "李四", Age: 21},
		{Name: "王五", Age: 22},
	}

	buildIndex(list)

	queryData("张三", 30)
}

func simpleHash(str string) (ret int) {
	for i := 0; i < len(str); i++ {
		// 取出字符
		c := str[i]
		// 将字符的ASCII码相加
		ret += int(c)
	}
	return
}

// 查询键
type classicQueryKey struct {
	Name string
	Age  int
}

func (c *classicQueryKey) hash() int {
	return simpleHash(c.Name) + c.Age*1000000
}

// 创建哈希值到数据的索引关系
var mp = make(map[int][]*Profile)

// 构建数据索引
func buildIndex(list []*Profile) {
	for _, profile := range list {

		// 构建数据的查询索引
		key := classicQueryKey{profile.Name, profile.Age}

		// 计算数据的哈希值，取出已经存在的记录
		existValue := mp[key.hash()]

		// 将当前数据添加到已经存在的记录中
		existValue = append(existValue, profile)

		// 将切片重新设置到映射中
		mp[key.hash()] = existValue
	}
}

func queryData(name string, age int) {
	// 根据给定查询条件构建查询键
	key2Query := classicQueryKey{name, age}

	// 计算查询键的哈希值并查询，获得相同的哈希值所有结果集合
	resultList := mp[key2Query.hash()]

	for _, result := range resultList {
		if result.Name == name && result.Age == age {
			fmt.Println(result)
			return
		}
	}

	fmt.Println("not found")
}
