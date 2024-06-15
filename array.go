package main

import (
	"fmt"
)

func main() {
	arr := [5]int{2, 24, 45, 7, 5}
	arr1 := [5]int{1: 25, 4: 99}
	arr[3] = 39
	arr1[2] = 88
	fmt.Println(arr)
	fmt.Println(arr1)
}
