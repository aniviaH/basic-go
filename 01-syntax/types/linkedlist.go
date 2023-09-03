package main

import "time"

/**
结构体定义

基本语法:
• type Name struct {
• fieldName FieldType
• // ...
• }

结构体和结构体的字段都遵循大小写控制访问性的原则。

通过 . 这个符号来访问字段或者方法。
*/

type LinkedList struct {
	head *node
	tail *node

	// 这个就是包外可以访问
	Len int

	CreateTime time.Time
}

func (l *LinkedList) Add(idx int, val any) error {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Append(val any) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) Delete(idx int) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LinkedList) toSlice() ([]any, error) {
	//TODO implement me
	panic("implement me")
}

//// 结构体接收器
//func (l LinkedList) Add(index int, val any) error {
//	// TODO implement me
//	return nil
//}
//
//// 指针接收器
//func (l *LinkedList) Add1(index int, val any) error {
//	// TODO implement me
//	return nil
//}

type node struct {
	prev *node
	next *node
}
