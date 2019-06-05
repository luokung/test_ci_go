package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello , test_ci_go")
	tbl := []int{1, 2, 3, 4}
	for i,v := range tbl {
		fmt.Println("i:", i, " v:", v)
	}
}