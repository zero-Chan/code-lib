package log

// 定义新日志格式至少按照如下标准

// 例:

// Create()函数返回该log实例
// func Create(...) MyLogger

// New()函数返回该log实例的指针
// func Create(...) *MyLogger

// Virtualize()行为函数返回 Logger interface
// func (this *MyLogger) Virtualize() Logger

// CreateDefault()函数返回通用的log实例
// func CreateDefault() MyLogger

// NewDefault()函数返回通用的log实例指针
// func CreateDefault() *MyLogger

// 使用

// 例:

// package main

//import (
//	"subassembly/log/golog"
//)

//func main() {
//	mylog := golog.NewDefault()
//	logger := mylog.Virtualize()
//	logger.Printf("Hello World")
//}
