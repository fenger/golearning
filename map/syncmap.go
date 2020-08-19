package _map

import (
	"fmt"
	"sync"
)

/**
sync.Map 有以下特性：
- 无须初始化，直接声明即可。
- sync.Map 不能使用 map 的方式进行取值和设置等操作，而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
- 使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，Range 参数中回调函数的返回值在需要继续迭代遍历时，返回 true，终止迭代遍历时，返回 false。
*/
func RunSyncMap() {
	var sence sync.Map

	sence.Store("greece", 97)
	sence.Store("london", 100)
	sence.Store("egypt", 200)

	fmt.Println(sence.Load("london"))

	sence.Delete("london")

	sence.Range(func(key, value interface{}) bool {
		fmt.Println("iterate: ", key, value)
		return true
	})
}
