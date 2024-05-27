package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	numbers := [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	map1 := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(map1)
	// 基于数组创建切片
	slice1 := numbers[1:3]
	if slice1 != nil {
		fmt.Println(slice1)
	}
	// 使用内置函数创建切片
	slice2 := make([]int, 2)
	slice2[1] = 199
	fmt.Println(slice2)
}
