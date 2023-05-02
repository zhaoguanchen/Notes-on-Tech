package main

import (
	"fmt"
	"strings"
)

func findIan() {
	var str string
	fmt.Printf("please enter string：")
	_, _ = fmt.Scan(&str)
	str = strings.ToLower(str)
	if strings.Contains(str, "a") && strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}

}
