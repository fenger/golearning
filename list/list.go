package _list

import (
	"container/list"
	"fmt"
)

func NewList() {
	l := list.New()
	// 尾部添加
	l.PushBack("first")
	// 头部添加
	l.PushFront(88)

	// 尾部添加后保存元素句柄
	element := l.PushBack("second")

	// 在second之后添加high
	l.InsertAfter("high", element)

	// 在second之前添加zero
	l.InsertBefore("zero", element)

	output(l)

	// 移除
	l.Remove(element)

	output(l)
}

func output(l *list.List) {
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Printf("%v ", i.Value)
	}
	fmt.Println()
}
