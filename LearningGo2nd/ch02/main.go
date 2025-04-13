package main

import (
	"fmt"
)

func main() {
	ex01()
	ex02()
	ex03()
}

func ex01() {
	var i int = 20
	var f float32 = float32(i)
	fmt.Print(i, f)
}

func ex02() {
	const value = 10
	var i int = value
	var f float32 = float32(value)
	fmt.Print(i, f)
}

func ex03() {
	var b byte = byte.MaxValue + 1
	var smallI int32 = int32.MaxValue + 1
	var bigI uint64 = uint64.MaxValue + 1
	fmt.Print(b, smallI, bigI)
}
