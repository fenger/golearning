package _nil

import (
	"fmt"
	"unsafe"
)

// nil不能比较
func CannotCompare() {
	// invalid operation: nil == nil (operator == not defined on nil)
	// fmt.Println(nil == nil)
}

// nil没有默认类型
func NoDefaultType() {
	fmt.Printf("%T", nil)
	// use of untyped nil
	// print(nil)
}

// 不同类型的nil指针是一样的
func SamePointNil() {
	var arr []int
	var num *int
	// 0x0
	fmt.Printf("%p\n", arr)
	// 0x0
	fmt.Printf("%p\n", num)
}

// 不同类型的nil是不能比较的
func DifferentTypeCannotCompare() {
	// var m map[int]string
	// var ptr *int
	// fmt.Printf(m == ptr)
}

// nil 是 map、slice、pointer、channel、func、interface 的零值
func ShowNilValue() {
	var m map[int]string
	var ptr *int
	var c chan int
	var sl []int
	// var f func()
	var i interface{}
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", ptr)
	fmt.Printf("%#v\n", c)
	fmt.Printf("%#v\n", sl)
	// fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", i)
}

// 不同类型的nil值占用的内存大小可能不一样
func ShowNilSize() {
	var p *struct{}
	// 8
	fmt.Println(unsafe.Sizeof(p))

	var s []int
	// 24
	fmt.Println(unsafe.Sizeof(s))

	var m map[int]string
	// 8
	fmt.Println(unsafe.Sizeof(m))

	var c chan string
	// 8
	fmt.Println(unsafe.Sizeof(c))

	var f func()
	// 8
	fmt.Println(unsafe.Sizeof(f))

	var i interface{}
	// 16
	fmt.Println(unsafe.Sizeof(i))
}
