package _new

import "fmt"

/**

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func new(Type) *Type
从上面的代码可以看出，new 函数只接受一个参数，这个参数是一个类型，并且返回一个指向该类型内存地址的指针。同时 new 函数会把分配的内存置为零，也就是类型的零值。

*/

func NewPrimary() {
	var sum *int
	sum = new(int)
	*sum = 98
	fmt.Println(*sum)
}

type Student struct {
	name string
	age  int
}

func NewStudent() {
	var s *Student
	s = new(Student)
	s.name = "zhangsan"

	fmt.Println(s)
}
