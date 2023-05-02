package main

import "fmt"

func trunc() {
	var num float32
	fmt.Printf("please enter floating point number：")
	_, _ = fmt.Scan(&num)
	var intValue = int32(num)
	fmt.Printf("the integer is : %v", intValue)
}
