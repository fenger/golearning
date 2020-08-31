package _make

import "fmt"

/**
make 也是用于内存分配的，但是和 new 不同，它只用于 chan、map 以及 slice 的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。

在Go语言中，make 函数的描述如下：
// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
// Slice: The size specifies the length. The capacity of the slice is
// equal to its length. A second integer argument may be provided to
// specify a different capacity; it must be no smaller than the
// length, so make([]int, 0, 10) allocates a slice of length 0 and
// capacity 10.
// Map: An empty map is allocated with enough space to hold the
// specified number of elements. The size may be omitted, in which case
// a small starting size is allocated.
// Channel: The channel's buffer is initialized with the specified
// buffer capacity. If zero, or the size is omitted, the channel is
// unbuffered.
func make(t Type, size ...IntegerType) Type
通过上面的代码可以看出 make 函数的 t 参数必须是 chan（通道）、map（字典）、slice（切片）中的一个，并且返回值也是类型本身。

注意：make 函数只用于 map，slice 和 channel，并且不返回指针。如果想要获得一个显式的指针，可以使用 new 函数进行分配，或者显式地使用一个变量的地址。

Go语言中的 new 和 make 主要区别如下：
1. make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
2. new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
3. new 分配的空间被清零。make 分配空间后，会进行初始化；

*/

func MakeMap() {
	var m map[string]int

	m = make(map[string]int)

	m["first"] = 0
	m["second"] = 1
	fmt.Println(m)
}

func MakeSlice() {
	var s []int
	s = make([]int, 2)
	s[0] = 1
	s[1] = 2
	// s[2] = 3
	fmt.Println(s)
}
