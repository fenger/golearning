package slice

import "fmt"

/**
从数组或切片生成新的切片拥有如下特性：
- 取出的元素数量为：结束位置 - 开始位置；
- 取出元素不包含结束位置对应的索引，切片最后一个元素使用 slice[len(slice)] 获取；
- 当缺省开始位置时，表示从连续区域开头到结束位置；
- 当缺省结束位置时，表示从开始位置到整个连续区域末尾；
- 两者同时缺省时，与切片本身等效；
- 两者同时为 0 时，等效于空切片，一般用于切片复位。

区分数组还是切片
1. 如果定义时指定长度，则为数组，如 a := [3]int{}
2. 如果定义时未指定长度，则为切片，如 a := []int{}

数组为值类型，切片为引用类型

*/
// generate slice
func GenerateSlice() {
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])
	// arr type: [3]int, slice type: []int
	fmt.Printf("arr type: %T, slice type: %T\n", a, a[1:2])

	b := []int{1, 2, 3}
	fmt.Printf("type b: %T\n", b)
}
